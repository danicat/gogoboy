package main

var opcodes = map[byte]struct {
	name   string
	cycles int
	exec   func(z *Z80)
}{
	0x00: {"NOP", 4, func(z *Z80) {}},

	0x3E: {"LD A, n", 8, func(z *Z80) { z.A = z.fetch() }},
	0x06: {"LD B, n", 8, func(z *Z80) { z.B = z.fetch() }},
	0x0E: {"LD C, n", 8, func(z *Z80) { z.C = z.fetch() }},
	0x16: {"LD D, n", 8, func(z *Z80) { z.D = z.fetch() }},
	0x1E: {"LD E, n", 8, func(z *Z80) { z.E = z.fetch() }},
	0x26: {"LD H, n", 8, func(z *Z80) { z.H = z.fetch() }},
	0x2E: {"LD L, n", 8, func(z *Z80) { z.L = z.fetch() }},

	0x7F: {"LD A, A", 4, func(z *Z80) {}},
	0x78: {"LD A, B", 4, func(z *Z80) { z.A = z.B }},
	0x79: {"LD A, C", 4, func(z *Z80) { z.A = z.C }},
	0x7A: {"LD A, D", 4, func(z *Z80) { z.A = z.D }},
	0x7B: {"LD A, E", 4, func(z *Z80) { z.A = z.E }},
	0x7C: {"LD A, H", 4, func(z *Z80) { z.A = z.H }},
	0x7D: {"LD A, L", 4, func(z *Z80) { z.A = z.L }},

	0x87: {"ADD A, A", 4, func(z *Z80) { z.A = z.add8(z.A, z.A) }},
	0x80: {"ADD A, B", 4, func(z *Z80) { z.A = z.add8(z.A, z.B) }},
	0x81: {"ADD A, C", 4, func(z *Z80) { z.A = z.add8(z.A, z.C) }},
	0x82: {"ADD A, D", 4, func(z *Z80) { z.A = z.add8(z.A, z.D) }},
	0x83: {"ADD A, E", 4, func(z *Z80) { z.A = z.add8(z.A, z.E) }},
	0x84: {"ADD A, H", 4, func(z *Z80) { z.A = z.add8(z.A, z.H) }},
	0x85: {"ADD A, L", 4, func(z *Z80) { z.A = z.add8(z.A, z.L) }},
}
