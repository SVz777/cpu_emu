package cpu

func NewTIS100(name string, hz Hz, size int) *CPU {
	cpu := CPU{
		Name: name,
		Hz:   hz,
		Regs: map[Rs]int64{
			ACC: 0,
			BAK: 0,
		},
		JumpTables: map[string]int64{},
		DataBuses: map[string]*DataBus{
			UP:    NewDataBus(size),
			DOWN:  NewDataBus(size),
			LEFT:  NewDataBus(size),
			RIGHT: NewDataBus(size),
		},
	}
	cpu.Instructions = map[string]func([]Addr){
		NOP: DoNOP(&cpu),
		MOV: DoMOV(&cpu),
		SAV: DoSAV(&cpu),
		SWP: DoSWP(&cpu),
		ADD: DoADD(&cpu),
		SUB: DoSUB(&cpu),
		JMP: DoJMP(&cpu),
		JEZ: DoJEZ(&cpu),
		JNZ: DoJNZ(&cpu),
		JGZ: DoJGZ(&cpu),
		JLZ: DoJLZ(&cpu),
		PRT: DoPRT(&cpu),
	}

	return &cpu
}
