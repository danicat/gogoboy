package cpu

func (z *Z80) SetZFlag() {
	z.F |= 0b10000000
}

func (z *Z80) ResetZFlag() {
	z.F &= 0b01111111
}

func (z *Z80) ZFlag() bool {
	return z.F&0b10000000 > 0
}

func (z *Z80) SetNFlag() {
	z.F |= 0b01000000
}

func (z *Z80) ResetNFlag() {
	z.F &= 0b10111111
}

func (z *Z80) NFlag() bool {
	return z.F&0b01000000 > 0
}

func (z *Z80) SetHFlag() {
	z.F |= 0b00100000
}

func (z *Z80) ResetHFlag() {
	z.F &= 0b11011111
}

func (z *Z80) HFlag() bool {
	return z.F&0b00100000 > 0
}

func (z *Z80) SetCFlag() {
	z.F |= 0b00010000
}

func (z *Z80) ResetCFlag() {
	z.F &= 0b11101111
}

func (z *Z80) CFlag() bool {
	return z.F&0b00010000 > 0
}
