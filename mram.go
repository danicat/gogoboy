package main

const MainRAMSize = 8192

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

func (m *MRAM) ReadAddr(a int16) byte {
	return m.ram[a]
}
