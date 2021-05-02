package main

// Z80 holds the internal representation of the Z80 CPU registers
type Z80 struct {
	PC  int16
	B   byte
	ram *MRAM
}

// NewZ80 creates a new Z80 instance that runs a given program
func NewZ80(program []byte) *Z80 {
	r := NewMRAM()
	r.LoadProgram(program)
	return &Z80{
		ram: r,
	}
}

// step runs one instruction at a time
func (z *Z80) step() {
	op := z.fetch()
	switch op {
	case 0x00: // NOP
	case 0x06: // LD B, n
		z.B = z.fetch()
	}
}

func (z *Z80) fetch() byte {
	op := z.ram.ReadAddr(z.PC)
	z.PC++
	return op
}
