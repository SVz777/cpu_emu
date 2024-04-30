package cpu

import (
	"time"
)

type Hz = time.Duration

const (
	GHz Hz = 1_000
	MHz Hz = 100
)

type Rs = string
type Ds = string

const (
	// reg

	ACC Rs = "ACC"
	BAK Rs = "BAK"

	// databus

	UP    Ds = "UP"
	LEFT  Ds = "LEFT"
	RIGHT Ds = "RIGHT"
	DOWN  Ds = "DOWN"
)

type As = string

const (
	NOP As = "NOP"
	MOV    = "MOV"
	SAV    = "SAV"
	SWP    = "SWP"
	ADD    = "ADD"
	SUB    = "SUB"
	JMP    = "JMP"
	JEZ    = "JEZ"
	JNZ    = "JNZ"
	JLZ    = "JLZ"
	JGZ    = "JGZ"
	PRT    = "PRT"
)

type AddrType uint8

const (
	TypeNone AddrType = iota
	AddrTypeConst
	AddrTypeTAG
	AddrTypeReg
	AddrTypeDataBus
)

type Addr struct {
	Typ   AddrType
	Name  string
	Const int64
}

type Prog struct {
	Line   int64
	Source string
	Tag    string
	Do     func([]Addr)
	Addrs  []Addr
}
