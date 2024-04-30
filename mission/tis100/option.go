package tis100

import (
	"github.com/SVz777/cpu_emu/cpu"
)

type Option func(m *Mission)

func WithCpus(cpus ...*cpu.CPU) Option {
	return func(m *Mission) {
		if m.cpus == nil {
			m.cpus = make(map[string]*cpu.CPU, len(cpus))
		}
		for _, c := range cpus {
			m.cpus[c.Name] = c
		}
	}
}

func WithSeq(inSeqs, outSeqs map[string][]int64) Option {
	return func(m *Mission) {
		m.inSeqs = inSeqs
		for k := range inSeqs {
			m.inDataBuses[k] = cpu.NewDataBus(1)
		}
		m.outSeqs = outSeqs
		for k := range outSeqs {
			m.outDataBuses[k] = cpu.NewDataBus(1)
		}
	}
}
