package cpu

func (z *Z80) AF() uint16 {
	return pair(z.A, z.F)
}

func (z *Z80) SetAF(v uint16) {
	writePair(&z.A, &z.F, v)
}

func (z *Z80) BC() uint16 {
	return pair(z.B, z.C)
}

func (z *Z80) SetBC(v uint16) {
	writePair(&z.B, &z.C, v)
}

func (z *Z80) DE() uint16 {
	return pair(z.D, z.E)
}

func (z *Z80) SetDE(v uint16) {
	writePair(&z.D, &z.E, v)
}

func (z *Z80) HL() uint16 {
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
