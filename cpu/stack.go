package cpu

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
