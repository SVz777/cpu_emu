package tis100

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SVz777/cpu_emu/cpu"
)

type Mission struct {
	Name         string
	inDataBuses  map[string]*cpu.DataBus
	outDataBuses map[string]*cpu.DataBus

	inSeqs  map[string][]int64
	outSeqs map[string][]int64
	cpus    map[string]*cpu.CPU
}

func NewMission(name string, opts ...Option) *Mission {
	m := &Mission{
		Name:         name,
		inDataBuses:  make(map[string]*cpu.DataBus),
		outDataBuses: make(map[string]*cpu.DataBus),
	}
	for _, opt := range opts {
		opt(m)
	}
	m.init()
	return m
}

func (m *Mission) init() {
	for _, c := range m.cpus {
		fileName := "./" + c.Name + ".svz"
		codes, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("read %s failed: %v", fileName, err)
		}
		prog, err := cpu.BuildProg(c, strings.Split(string(codes), "\n"))
		if err != nil {
			panic(err)
		}
		c.LoadProg(prog)
	}
}

func (m *Mission) Run() bool {
	for _, c := range m.cpus {
		go func(cc *cpu.CPU) {
			for {
				cc.Run()
				cc.ResetPC()
			}
		}(c)
	}

	for name, in := range m.inSeqs {
		go func(name string, ins []int64) {
			for _, v := range ins {
				m.inDataBuses[name].Write(v)
			}
		}(name, in)
	}
	var flag bool
	for name, out := range m.outSeqs {
		for _, v := range out {
			read := m.outDataBuses[name].Read()
			if v != read {
				flag = true
			}
			fmt.Printf("%10s | equal: %v, outSeq: %d, read: %d\n", name, v == read, v, read)
		}
	}
	return flag
}

func (m *Mission) In(cpuName string, ds cpu.Ds) {
	m.inDataBuses["in_"+cpuName].Connect(m.cpus[cpuName].DataBus(ds))
}

func (m *Mission) Out(cpuName string, ds cpu.Ds) {
	m.outDataBuses["out_"+cpuName].Connect(m.cpus[cpuName].DataBus(ds))
}
