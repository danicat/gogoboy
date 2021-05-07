package cpu

func (z *Z80) call(addr uint16) {
	z.push(z.PC + 1)
	z.jump(addr)
}

func (z *Z80) jump(addr uint16) {
	z.PC = addr
}
