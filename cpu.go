package main

// Z80 holds the internal representation of the Z80 CPU registers
type Z80 struct {
	PC  int16
	ram [1]byte
}

// NewZ80 creates a new Z80 instance that runs a given program
func NewZ80(program [1]byte) *Z80 {
	return &Z80{
		ram: program,
	}
}

func (z *Z80) fetch() byte {
	op := z.ram[z.PC]
	z.PC++
	return op
}
