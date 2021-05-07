package memory

const MemorySize = 65536

type Memory struct {
	data [MemorySize]byte
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) LoadProgram(p []byte, addr uint16) {
	var data [MemorySize]byte
	for i, b := range p {
		data[i+int(addr)] = b
	}
	m.data = data
}

func (m *Memory) Read(addr uint16) byte {
	return m.data[addr]
}

func (m *Memory) Write(addr uint16, val byte) {
	m.data[addr] = val
}

func (m *Memory) Write16(addr uint16, val uint16) {
	hi, lo := split(val)
	m.data[addr] = hi
	m.data[addr+1] = lo
}

func split(v uint16) (hi, lo byte) {
	hi = byte(v / 0x100)
	lo = byte(v % 0x100)
	return
}
