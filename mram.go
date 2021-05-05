package main

const MainRAMSize = 65536

type MRAM struct {
	ram [MainRAMSize]byte
}

func NewMRAM() *MRAM {
	return &MRAM{}
}

func (m *MRAM) LoadProgram(p []byte) {
	var ram [MainRAMSize]byte
	for i, b := range p {
		ram[i] = b
	}
	m.ram = ram
}

func (m *MRAM) ReadAddr8(hi, lo byte) byte {
	a := uint16(hi)*0x100 + uint16(lo)
	return m.ReadAddr(a)
}

func (m *MRAM) ReadAddr(a uint16) byte {
	return m.ram[a]
}
