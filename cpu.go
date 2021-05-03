package main

// Z80 holds the internal representation of the Z80 CPU registers
type Z80 struct {
	PC                     int16
	A, F, B, C, D, E, H, L byte
	ram                    *MRAM
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
	z.F = 0
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

	// 8-bit load
	case 0x06:
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

	// 8-bit ALU
	case 0x87: // ADD A
		z.A = z.add8(z.A, z.A)
	case 0x80: // ADD B
		z.A = z.add8(z.A, z.B)
	case 0x81: // ADD C
		z.A = z.add8(z.A, z.C)
	case 0x82: // ADD D
		z.A = z.add8(z.A, z.D)
	case 0x83: // ADD E
		z.A = z.add8(z.A, z.E)
	case 0x84: // ADD H
		z.A = z.add8(z.A, z.H)
	case 0x85: // ADD L
		z.A = z.add8(z.A, z.L)
	}

}

func (z *Z80) add8(l, r byte) byte {
	var res int16 = int16(l) + int16(r)
	if res == 0 {
		z.SetZFlag()
	}
	z.ResetNFlag()
	if l&0x0F+r&0x0F > 0x0F {
		z.SetHFlag()
	}
	if res > 0xFF {
		z.SetCFlag()
	}
	return byte(res)
}

func (z *Z80) SetZFlag() {
	z.F |= 0b10000000
}

func (z *Z80) ResetZFlag() {
	z.F &= 0b01111111
}

func (z *Z80) ZFlag() bool {
	return z.F&0b10000000 > 0
}

func (z *Z80) SetNFlag() {
	z.F |= 0b01000000
}

func (z *Z80) ResetNFlag() {
	z.F &= 0b10111111
}

func (z *Z80) NFlag() bool {
	return z.F&0b01000000 > 0
}

func (z *Z80) SetHFlag() {
	z.F |= 0b00100000
}

func (z *Z80) ResetHFlag() {
	z.F &= 0b11011111
}

func (z *Z80) HFlag() bool {
	return z.F&0b00100000 > 0
}

func (z *Z80) SetCFlag() {
	z.F |= 0b00010000
}

func (z *Z80) ResetCFlag() {
	z.F &= 0b11101111
}

func (z *Z80) CFlag() bool {
	return z.F&0b00010000 > 0
}

func (z *Z80) fetch() byte {
	op := z.ram.ReadAddr(z.PC)
	z.PC++
	return op
}
