package cpu

func (z *Z80) call(flag bool) {
	hi := z.fetch()
	lo := z.fetch()
	addr := pair(hi, lo)

	if flag {
		z.push(z.PC + 1)
		z.PC = addr
	}
}

// func (z *Z80) jump(addr uint16) {
// 	z.PC = addr
// }
