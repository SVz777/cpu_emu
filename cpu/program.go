package cpu

import (
	"fmt"
	"strconv"
	"strings"
)

func BuildProg(cpu *CPU, codes []string) ([]*Prog, error) {
	progs := make([]*Prog, 0, len(codes))
	for idx := 0; idx < len(codes); idx++ {
		code := strings.TrimSpace(codes[idx])
		if code == "" {
			continue
		}
		prog := &Prog{
			Line:   int64(idx),
			Source: code,
		}
		// 判断是不是跳转标签
		if strings.HasSuffix(code, ":") {
			prog.Tag = code[:len(code)-1]
			prog.Do = cpu.Instructions[NOP]
		} else {
			line := strings.SplitN(code, " ", 2)
			if doFunc, ok := cpu.Instructions[strings.TrimSpace(line[0])]; ok {
				prog.Do = doFunc
			} else {
				cpu.Panic(fmt.Sprintf("unhandled instruction: %v", line[0]))
			}
			if len(line) == 2 {
				addrs := strings.Split(strings.TrimSpace(line[1]), ",")
				for _, addr := range addrs {
					addr = strings.TrimSpace(addr)
					if len(addr) < 2 {
						return nil, fmt.Errorf("invalid addr: %v", addr)
					}
					a := Addr{}
					if addr[0] == '$' {
						v, err := strconv.ParseInt(addr[1:], 10, 64)
						if err != nil {
							return nil, fmt.Errorf("invalid addr: %v", addr)
						}
						a.Typ = AddrTypeConst
						a.Const = v
					} else if _, ok := cpu.Regs[addr]; ok { // addr{
						a.Typ = AddrTypeReg
						a.Name = addr
					} else if _, ok := cpu.DataBuses[addr]; ok {
						a.Typ = AddrTypeDataBus
						a.Name = addr
					} else {
						a.Typ = AddrTypeTAG
						a.Name = addr
					}
					prog.Addrs = append(prog.Addrs, a)
				}
			}
		}
		progs = append(progs, prog)
	}
	return progs, nil
}
