package cpu

var opcodes = map[byte]struct {
	name   string
	cycles int
	exec   func(z *Z80)
}{
	0x00: {"NOP", 4, func(z *Z80) {}},

	// 8-Bit Loads
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

	0x66: {"LD H, (HL)", 8, func(z *Z80) { z.H = z.ram.Read(z.HL()) }},

	// 16-Bit Loads
	0x01: {"LD BC, nn", 12, func(z *Z80) { z.B = z.fetch(); z.C = z.fetch() }},
	0x11: {"LD DE, nn", 12, func(z *Z80) { z.D = z.fetch(); z.E = z.fetch() }},
	0x21: {"LD HL, nn", 12, func(z *Z80) { z.H = z.fetch(); z.L = z.fetch() }},
	0x31: {"LD SP, nn", 12, func(z *Z80) { hi := z.fetch(); lo := z.fetch(); z.SP = pair(hi, lo) }},

	// 8-Bit ALU
	0x87: {"ADD A, A", 4, func(z *Z80) { z.A = z.add8(z.A, z.A, false) }},
	0x80: {"ADD A, B", 4, func(z *Z80) { z.A = z.add8(z.A, z.B, false) }},
	0x81: {"ADD A, C", 4, func(z *Z80) { z.A = z.add8(z.A, z.C, false) }},
	0x82: {"ADD A, D", 4, func(z *Z80) { z.A = z.add8(z.A, z.D, false) }},
	0x83: {"ADD A, E", 4, func(z *Z80) { z.A = z.add8(z.A, z.E, false) }},
	0x84: {"ADD A, H", 4, func(z *Z80) { z.A = z.add8(z.A, z.H, false) }},
	0x85: {"ADD A, L", 4, func(z *Z80) { z.A = z.add8(z.A, z.L, false) }},
	0xC6: {"ADC A, #", 8, func(z *Z80) { z.A = z.add8(z.A, z.fetch(), false) }},

	0x8F: {"ADC A, A", 4, func(z *Z80) { z.A = z.add8(z.A, z.A, z.CFlag()) }},
	0x88: {"ADC A, B", 4, func(z *Z80) { z.A = z.add8(z.A, z.B, z.CFlag()) }},
	0x89: {"ADC A, C", 4, func(z *Z80) { z.A = z.add8(z.A, z.C, z.CFlag()) }},
	0x8A: {"ADC A, D", 4, func(z *Z80) { z.A = z.add8(z.A, z.D, z.CFlag()) }},
	0x8B: {"ADC A, E", 4, func(z *Z80) { z.A = z.add8(z.A, z.E, z.CFlag()) }},
	0x8C: {"ADC A, H", 4, func(z *Z80) { z.A = z.add8(z.A, z.H, z.CFlag()) }},
	0x8D: {"ADC A, L", 4, func(z *Z80) { z.A = z.add8(z.A, z.L, z.CFlag()) }},
	0xCE: {"ADC A, #", 8, func(z *Z80) { z.A = z.add8(z.A, z.fetch(), z.CFlag()) }},

	0x03: {"INC BC", 8, func(z *Z80) { z.inc(&z.B, &z.C) }},
	0x13: {"INC DE", 8, func(z *Z80) { z.inc(&z.D, &z.E) }},
	0x23: {"INC HL", 8, func(z *Z80) { z.inc(&z.H, &z.L) }},
	0x33: {"INC SP", 8, func(z *Z80) { z.SP++ }},

	0x0B: {"DEC BC", 8, func(z *Z80) { z.dec(&z.B, &z.C) }},
	0x1B: {"DEC DE", 8, func(z *Z80) { z.dec(&z.D, &z.E) }},
	0x2B: {"DEC HL", 8, func(z *Z80) { z.dec(&z.H, &z.L) }},
	0x3B: {"DEC SP", 8, func(z *Z80) { z.SP-- }},

	// Flow control
	0xCC: {"CALL Z, nn", 12, func(z *Z80) { z.call(z.ZFlag()) }},
	0xC4: {"CALL NZ, nn", 12, func(z *Z80) { z.call(!z.ZFlag()) }},
	0xDC: {"CALL C, nn", 12, func(z *Z80) { z.call(z.CFlag()) }},
	0xD4: {"CALL NC, nn", 12, func(z *Z80) { z.call(!z.CFlag()) }},

	0xF5: {"PUSH AF", 16, func(z *Z80) { z.push(z.AF()) }},
	0xC5: {"PUSH BC", 16, func(z *Z80) { z.push(z.BC()) }},
	0xD5: {"PUSH DE", 16, func(z *Z80) { z.push(z.DE()) }},
	0xE5: {"PUSH HL", 16, func(z *Z80) { z.push(z.HL()) }},

	0xF1: {"POP AF", 16, func(z *Z80) { z.SetAF(z.pop()) }},
	0xC1: {"POP BC", 16, func(z *Z80) { z.SetBC(z.pop()) }},
	0xD1: {"POP DE", 16, func(z *Z80) { z.SetDE(z.pop()) }},
	0xE1: {"POP HL", 16, func(z *Z80) { z.SetHL(z.pop()) }},
}
