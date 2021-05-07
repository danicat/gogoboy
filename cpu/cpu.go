package cpu

import (
	"fmt"

	"github.com/danicat/gogoboy/memory"
)

// Z80 holds the internal representation of the Z80 CPU registers
type Z80 struct {
	PC, SP                 uint16
	A, F, B, C, D, E, H, L byte
	ram                    *memory.Memory
	cycles                 int
	maxCycles              int
}

// NewZ80 creates a new Z80 instance
func NewZ80() *Z80 {
	r := memory.NewMemory()
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

func (z *Z80) fetch() byte {
	op := z.ram.Read(z.PC)
	z.PC++
	return op
}
