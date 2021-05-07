package cpu

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
