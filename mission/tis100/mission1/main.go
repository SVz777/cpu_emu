package main

import (
	"fmt"

	"github.com/SVz777/cpu_emu/cpu"
	"github.com/SVz777/cpu_emu/mission/tis100"
)

// 2*a
// in
//
//	|
//	v
//
// c0
//
//	|
//	v
//
// out
func main() {
	var (
		inSeq  = map[string][]int64{"in_cpu0": {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
		outSeq = map[string][]int64{"out_cpu0": {2, 4, 6, 8, 10, 12, 14, 16, 18, 20}}
	)
	m := tis100.NewMission("2*a",
		tis100.WithSeq(inSeq, outSeq),
		tis100.WithCpus(cpu.NewTIS100("cpu0", cpu.MHz, 0)),
	)
	m.In("cpu0", cpu.UP)
	m.Out("cpu0", cpu.DOWN)
	flag := m.Run()
	if flag {
		fmt.Printf("mission %s failed\n", m.Name)
	} else {
		fmt.Printf("mission %s completed\n", m.Name)
	}
}
