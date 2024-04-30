package cpu

import (
	"fmt"
	"runtime"
	"time"
)

type DataBus struct {
	in  chan int64
	out chan int64
}

func (d *DataBus) In() chan int64 {
	return d.in
}
func (d *DataBus) Out() chan int64 {
	return d.out
}

func (d *DataBus) Write(v int64) {
	d.out <- v
}

func (d *DataBus) Read() int64 {
	return <-d.in
}

func (d *DataBus) Connect(other *DataBus) {
	d.in = other.out
	d.out = other.in
}

func NewDataBus(size int) *DataBus {
	return &DataBus{
		in:  make(chan int64, size),
		out: make(chan int64, size),
	}
}

type CPU struct {
	Name  string
	Hz    Hz
	Progs []*Prog

	Instructions map[As]func([]Addr)
	Regs         map[Rs]int64

	JumpTables map[string]int64
	PC         int64
	NowPC      int64
	DataBuses  map[string]*DataBus
}

func (cpu *CPU) Connect(ds Ds, other *CPU, otherDs Ds) {
	cpu.DataBus(ds).Connect(other.DataBus(otherDs))
}

func (cpu *CPU) Reset() {
	for k := range cpu.Regs {
		cpu.Regs[k] = 0
	}
	cpu.ResetPC()
}

func (cpu *CPU) ResetPC() {
	cpu.PC = 0
	cpu.NowPC = 0
}

func (cpu *CPU) Stat() {
	fmt.Printf("PC: %d\t|%-15s\t|%5v|\n", cpu.NowPC, cpu.Progs[cpu.NowPC].Source, cpu.Regs)
}

func (cpu *CPU) Panic(reason string) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println(cpu.Name, " panic: ", reason)
	cpu.Stat()
	fmt.Printf("goroutine err stack:%s\n", string(buf[:n]))
}

func (cpu *CPU) Run() {
	for {
		as := cpu.loadAs()
		if as == nil {
			return
		}
		cpu.Do(as)
		//cpu.Stat()
		time.Sleep(time.Second / cpu.Hz)
	}
}

// ------------变量获取---------------

func (cpu *CPU) Reg(name string) int64 {
	reg, ok := cpu.Regs[name]
	if !ok {
		cpu.Panic(fmt.Sprintf("reg not found: %v", name))
	}
	return reg
}

func (cpu *CPU) DataBus(name string) *DataBus {
	databus, ok := cpu.DataBuses[name]
	if !ok {
		cpu.Panic(fmt.Sprintf("reg not found: %v", name))
	}
	return databus
}

// ------------数据读写---------------

func (cpu *CPU) WriteReg(name string, v int64) {
	if _, ok := cpu.Regs[name]; ok {
		cpu.Regs[name] = v
	} else {
		cpu.Panic(fmt.Sprintf("reg not found: %v", name))
	}
}

func (cpu *CPU) Recv(addr Addr) int64 {
	switch addr.Typ {
	case AddrTypeConst:
		return addr.Const
	case AddrTypeReg:
		return cpu.Reg(addr.Name)
	case AddrTypeDataBus:
		return cpu.DataBus(addr.Name).Read()
	default:
		cpu.Panic(fmt.Sprintf("unhandled type: %v", addr.Typ))
	}
	return 0
}

func (cpu *CPU) Send(addr Addr, v int64) {
	switch addr.Typ {
	case AddrTypeConst:
		cpu.Panic("const not support send")
	case AddrTypeReg:
		cpu.WriteReg(addr.Name, v)
	case AddrTypeDataBus:
		cpu.DataBus(addr.Name).Write(v)
	default:
		cpu.Panic(fmt.Sprintf("unhandled type: %v", addr.Typ))
	}
}

// ------------程序加载---------------

func (cpu *CPU) Do(prog *Prog) {
	prog.Do(prog.Addrs)
}

func (cpu *CPU) LoadProg(progs []*Prog) {
	cpu.Progs = progs
	for idx, prog := range progs {
		if prog.Tag != "" {
			cpu.JumpTables[prog.Tag] = int64(idx)
		}
	}
}

func (cpu *CPU) loadAs() *Prog {
	now := cpu.PC
	cpu.NowPC = now
	cpu.WritePC(now + 1)

	if int(now) >= len(cpu.Progs) {
		return nil
	}
	return cpu.Progs[now]
}

func (cpu *CPU) WritePC(pc int64) {
	cpu.PC = pc
}
