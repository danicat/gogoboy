package main

var opcodes = map[byte]struct {
	name   string
	cycles int
}{
	00: {"NOP", 4},

	0x3E: {"LD A, n", 8},
	0x06: {"LD B, n", 8},
	0x0E: {"LD C, n", 8},
	0x16: {"LD D, n", 8},
	0x1E: {"LD E, n", 8},
	0x26: {"LD H, n", 8},
	0x2E: {"LD L, n", 8},

	0x7F: {"LD A, A", 4},
	0x78: {"LD A, B", 4},
	0x79: {"LD A, C", 4},
	0x7A: {"LD A, D", 4},
	0x7B: {"LD A, E", 4},
	0x7C: {"LD A, H", 4},
	0x7D: {"LD A, L", 4},

	0x87: {"ADD A, A", 4},
	0x80: {"ADD A, B", 4},
	0x81: {"ADD A, C", 4},
	0x82: {"ADD A, D", 4},
	0x83: {"ADD A, E", 4},
	0x84: {"ADD A, H", 4},
	0x85: {"ADD A, L", 4},
}
