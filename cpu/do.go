package cpu

import (
	"fmt"
)

func DoNOP(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {

	}
}

func DoMOV(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 2 {
			panic(fmt.Sprintf("unhandled MOV: %v", addrs))
		}
		fromAddr, toAddr := addrs[0], addrs[1]
		cpu.Send(toAddr, cpu.Recv(fromAddr))
	}
}

func DoSAV(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 0 {
			panic(fmt.Sprintf("unhandled SAV: %v", addrs))
		}
		cpu.WriteReg(BAK, cpu.Reg(ACC))
	}
}

func DoSWP(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 0 {
			panic(fmt.Sprintf("unhandled SWP: %v", addrs))
		}
		bak := cpu.Reg(BAK)
		acc := cpu.Reg(ACC)
		cpu.WriteReg(ACC, bak)
		cpu.WriteReg(BAK, acc)
	}
}

func DoADD(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled ADD: %v", addrs))
		}
		fromAddr := addrs[0]
		cpu.WriteReg(ACC, cpu.Reg(ACC)+cpu.Recv(fromAddr))
	}
}

func DoSUB(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled SUB: %v", addrs))
		}
		fromAddr := addrs[0]
		cpu.WriteReg(ACC, cpu.Reg(ACC)-cpu.Recv(fromAddr))
	}
}

func DoJMP(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled JMP: %v", addrs))
		}
		toAddr := addrs[0]
		if toAddr.Typ != AddrTypeTAG {
			panic(fmt.Sprintf("JMP: toAddr is not TAG: %v", toAddr))
		}
		var flag bool
		for tag, pc := range cpu.JumpTables {
			if tag == toAddr.Name {
				flag = true
				cpu.WritePC(pc)
				break
			}
		}
		if !flag {
			panic(fmt.Sprintf("JMP: not found: %v", toAddr.Name))
		}
	}
}

func DoJEZ(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled JEZ: %v", addrs))
		}
		toAddr := addrs[0]
		if toAddr.Typ != AddrTypeTAG {
			panic(fmt.Sprintf("JEZ: toAddr is not TAG: %v", toAddr))
		}
		if cpu.Reg(ACC) != 0 {
			return
		}
		var flag bool
		for tag, pc := range cpu.JumpTables {
			if tag == toAddr.Name {
				flag = true
				cpu.WritePC(pc)
				break
			}
		}
		if !flag {
			panic(fmt.Sprintf("JEZ: not found: %v", toAddr.Name))
		}
	}
}

func DoJNZ(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled JNZ: %v", addrs))
		}
		toAddr := addrs[0]
		if toAddr.Typ != AddrTypeTAG {
			panic(fmt.Sprintf("JNZ: toAddr is not TAG: %v", toAddr))
		}
		if cpu.Reg(ACC) == 0 {
			return
		}
		var flag bool
		for tag, pc := range cpu.JumpTables {
			if tag == toAddr.Name {
				flag = true
				cpu.WritePC(pc)
				break
			}
		}
		if !flag {
			panic(fmt.Sprintf("JNZ: not found: %v", toAddr.Name))
		}
	}
}

func DoJGZ(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled JGZ: %v", addrs))
		}
		toAddr := addrs[0]
		if toAddr.Typ != AddrTypeTAG {
			panic(fmt.Sprintf("JGZ: toAddr is not TAG: %v", toAddr))
		}
		if cpu.Reg(ACC) < 0 {
			return
		}
		var flag bool
		for tag, pc := range cpu.JumpTables {
			if tag == toAddr.Name {
				flag = true
				cpu.WritePC(pc)
				break
			}
		}
		if !flag {
			panic(fmt.Sprintf("JGZ: not found: %v", toAddr.Name))
		}
	}
}

func DoJLZ(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled JLZ: %v", addrs))
		}
		toAddr := addrs[0]
		if toAddr.Typ != AddrTypeTAG {
			panic(fmt.Sprintf("JLZ: toAddr is not TAG: %v", toAddr))
		}
		if cpu.Reg(ACC) > 0 {
			return
		}
		var flag bool
		for tag, pc := range cpu.JumpTables {
			if tag == toAddr.Name {
				flag = true
				cpu.WritePC(pc)
				break
			}
		}
		if !flag {
			panic(fmt.Sprintf("JLZ: not found: %v", toAddr.Name))
		}
	}
}

func DoPRT(cpu *CPU) func([]Addr) {
	return func(addrs []Addr) {
		if len(addrs) != 1 {
			panic(fmt.Sprintf("unhandled PRT: %v", addrs))
		}
		fromAddr := addrs[0]
		fmt.Println(cpu.Name, ":", cpu.Recv(fromAddr))
	}
}
