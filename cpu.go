package main

import (
	"fmt"
)

// Z80 holds the internal representation of the Z80 CPU registers
type Z80 struct {
	PC, SP                 uint16
	A, F, B, C, D, E, H, L byte
	ram                    *MRAM
	cycles                 int
	maxCycles              int
}

// NewZ80 creates a new Z80 instance
func NewZ80() *Z80 {
	r := NewMRAM()
	return &Z80{
		A:   0x01,
		F:   0xB0,
		B:   0x00,
		C:   0x13,
		D:   0x00,
		E:   0xD8,
		H:   0x01,
		L:   0x4D,
		PC:  0x100,
		SP:  0xFFFE,
		ram: r,
	}
}

// SetMaxCycles set the maximum number of cycles for the Run function to process. Max cycles of 0 indicates no limit.
func (z *Z80) SetMaxCycles(max int) {
	z.maxCycles = max
}

func (z *Z80) LoadProgram(p []byte, addr uint16) {
	z.PC = addr
	z.ram.LoadProgram(p, addr)
}

func (z *Z80) Reset() {
	z.PC = 0
	z.SP = 0xFFFE
	z.A = 0
	z.F = 0
	z.B = 0
	z.C = 0
	z.D = 0
	z.E = 0
	z.H = 0
	z.L = 0
}

func (z *Z80) Run() error {
	for z.cycles < z.maxCycles || z.maxCycles == 0 {
		err := z.step()
		if err != nil {
			return err
		}
	}
	return nil
}

// step runs one instruction at a time
func (z *Z80) step() error {
	op := z.fetch()

	inst, ok := opcodes[op]
	if !ok {
		return fmt.Errorf("opcode not implemented: %X", op)
	}

	z.cycles += inst.cycles
	inst.exec(z)

	return nil
}

func (z *Z80) add8(l, r byte, carry bool) byte {
	var res int16 = int16(l) + int16(r)
	var halfCarry = l&0x0F + r&0x0F

	if carry {
		res++
		halfCarry++
	}

	if res == 0 {
		z.SetZFlag()
	} else {
		z.ResetZFlag()
	}

	z.ResetNFlag()

	if halfCarry > 0x0F {
		z.SetHFlag()
	} else {
		z.ResetHFlag()
	}

	if res > 0xFF {
		z.SetCFlag()
	} else {
		z.ResetCFlag()
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
	op := z.ram.Read(z.PC)
	z.PC++
	return op
}

func (z *Z80) GetAF() uint16 {
	return pair(z.A, z.F)
}

func (z *Z80) SetAF(v uint16) {
	writePair(&z.A, &z.F, v)
}

func (z *Z80) GetBC() uint16 {
	return pair(z.B, z.C)
}

func (z *Z80) SetBC(v uint16) {
	writePair(&z.B, &z.C, v)
}

func (z *Z80) GetDE() uint16 {
	return pair(z.D, z.E)
}

func (z *Z80) SetDE(v uint16) {
	writePair(&z.D, &z.E, v)
}

func (z *Z80) GetHL() uint16 {
	return pair(z.H, z.L)
}

func (z *Z80) SetHL(v uint16) {
	writePair(&z.H, &z.L, v)
}

func pair(hi, lo byte) uint16 {
	return uint16(hi)*0x100 + uint16(lo)
}

func writePair(hi, lo *byte, v uint16) {
	*hi = byte(v / 0x100)
	*lo = byte(v % 0x100)
}

func split(v uint16) (hi, lo byte) {
	hi = byte(v / 0x100)
	lo = byte(v % 0x100)
	return
}

func (z *Z80) push(v uint16) {
	hi, lo := split(v)
	z.SP--
	z.ram.Write(z.SP, lo)
	z.SP--
	z.ram.Write(z.SP, hi)
}

func (z *Z80) pop() uint16 {
	hi := z.ram.Read(z.SP)
	z.SP++
	lo := z.ram.Read(z.SP)
	z.SP++
	return pair(hi, lo)
}

func (z *Z80) call(addr uint16) {
	z.push(z.PC + 1)
	z.jump(addr)
}

func (z *Z80) jump(addr uint16) {
	z.PC = addr
}
