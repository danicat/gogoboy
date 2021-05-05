package main

const MainRAMSize = 65536

type MRAM struct {
	ram [MainRAMSize]byte
}

func NewMRAM() *MRAM {
	return &MRAM{}
}

func (m *MRAM) LoadProgram(p []byte, addr uint16) {
	var ram [MainRAMSize]byte
	for i, b := range p {
		ram[i+int(addr)] = b
	}
	m.ram = ram
}

func (m *MRAM) Read(addr uint16) byte {
	return m.ram[addr]
}

func (m *MRAM) Write(addr uint16, val byte) {
	m.ram[addr] = val
}
