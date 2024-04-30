package main

import (
	"fmt"

	"github.com/SVz777/cpu_emu/cpu"
	"github.com/SVz777/cpu_emu/mission/tis100"
)

// a+b
// in   in
//
//	|    |
//	v    v
//
// c0   c1
//
//	|    |
//	v    |
//
// c2 <---
//
//	|
//	v
//
// out
func main() {
	var (
		inSeq = map[string][]int64{
			"in_cpu0": {0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			"in_cpu1": {1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
		}
		outSeq = map[string][]int64{
			"out_cpu2": {1, 5, 9, 13, 17, 21, 25, 29, 33, 37},
		}
	)
	/*

	 */
	cpus := []*cpu.CPU{
		cpu.NewTIS100("cpu0", cpu.MHz, 0),
		cpu.NewTIS100("cpu1", cpu.MHz, 0),
		cpu.NewTIS100("cpu2", cpu.MHz, 0),
	}
	cpus[2].Connect(cpu.UP, cpus[0], cpu.DOWN)
	cpus[2].Connect(cpu.RIGHT, cpus[1], cpu.DOWN)

	m := tis100.NewMission("a+b",
		tis100.WithSeq(inSeq, outSeq),
		tis100.WithCpus(cpus...),
	)
	m.In("cpu0", cpu.UP)
	m.In("cpu1", cpu.UP)
	m.Out("cpu2", cpu.DOWN)
	flag := m.Run()
	if flag {
		fmt.Printf("mission %s failed\n", m.Name)
	} else {
		fmt.Printf("mission %s completed\n", m.Name)
	}
}
