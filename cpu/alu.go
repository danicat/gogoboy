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

func (z *Z80) dec(hi, lo *byte) {
	val := pair(*hi, *lo)
	val--
	*hi, *lo = split(val)
}
