package main

// Z80 holds the internal representation of the Z80 CPU registers
type Z80 struct {
	PC                  int16
	A, B, C, D, E, H, L byte
	ram                 *MRAM
}

// NewZ80 creates a new Z80 instance that runs a given program
func NewZ80(program []byte) *Z80 {
	r := NewMRAM()
	r.LoadProgram(program)
	return &Z80{
		ram: r,
	}
}

func (z *Z80) LoadProgram(p []byte) {
	z.ram.LoadProgram(p)
	z.PC = 0
}

func (z *Z80) Reset() {
	z.PC = 0
	z.A = 0
	z.B = 0
	z.C = 0
	z.D = 0
	z.E = 0
	z.H = 0
	z.L = 0
}

// step runs one instruction at a time
func (z *Z80) step() {
	op := z.fetch()
	switch op {
	case 0x00: // NOP
	case 0x06: // LD B, n
		z.B = z.fetch()
	case 0x0E: // LD C, n
		z.C = z.fetch()
	case 0x16: // LD D, n
		z.D = z.fetch()
	case 0x1E: // LD E, n
		z.E = z.fetch()
	case 0x26: // LD H, n
		z.H = z.fetch()
	case 0x2E: // LD L, n
		z.L = z.fetch()
	case 0x7F: // LD A, A
	case 0x78: // LD A, B
		z.A = z.B
	case 0x79: // LD A, C
		z.A = z.C
	case 0x7A: // LD A, D
		z.A = z.D
	case 0x7B: // LD A, E
		z.A = z.E
	case 0x7C: // LD A, H
		z.A = z.H
	case 0x7D: // LD A, L
		z.A = z.L
	}
}

func (z *Z80) fetch() byte {
	op := z.ram.ReadAddr(z.PC)
	z.PC++
	return op
}
