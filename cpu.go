package main

import "log"

// Z80 holds the internal representation of the Z80 CPU registers
type Z80 struct {
	PC                     int16
	A, F, B, C, D, E, H, L byte
	ram                    *MRAM
	cycles                 int
	maxCycles              int
}

// NewZ80 creates a new Z80 instance
func NewZ80() *Z80 {
	r := NewMRAM()
	return &Z80{
		ram: r,
	}
}

// SetMaxCycles set the maximum number of cycles for the Run function to process. Max cycles of 0 indicates no limit.
func (z *Z80) SetMaxCycles(max int) {
	z.maxCycles = max
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

func (z *Z80) Run() {
	for z.cycles < z.maxCycles || z.maxCycles == 0 {
		z.step()
	}
}

// step runs one instruction at a time
func (z *Z80) step() {
	op := z.fetch()

	inst, ok := opcodes[op]
	if !ok {
		log.Fatalf("opcode not implemented: %x", op)
	}

	z.cycles += inst.cycles
	inst.exec(z)
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
