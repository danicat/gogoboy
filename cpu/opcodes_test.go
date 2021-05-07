package cpu

import (
	"testing"

	"github.com/danicat/gogoboy/memory"
)

func TestNOP(t *testing.T) {
	input := []byte{0, 0}
	expected := uint16(2)
	z := NewZ80()
	z.LoadProgram(input, 0)
	z.step()
	z.step()
	if z.PC != expected {
		t.Fatalf("expected %d, got %d", expected, z.PC)
	}
}

func TestLD(t *testing.T) {
	z := NewZ80()
	tbl := []struct {
		name     string
		program  []byte
		register *byte
		expected byte
	}{
		{
			"LD A, n",
			[]byte{0x3E, 0xAA},
			&z.A,
			0xAA,
		},
		{
			"LD B, n",
			[]byte{0x06, 0xDE},
			&z.B,
			0xDE,
		},
		{
			"LD C, n",
			[]byte{0x0E, 0xAD},
			&z.C,
			0xAD,
		},
		{
			"LD D, n",
			[]byte{0x16, 0xBE},
			&z.D,
			0xBE,
		},
		{
			"LD E, n",
			[]byte{0x1E, 0xEF},
			&z.E,
			0xEF,
		},
		{
			"LD H, n",
			[]byte{0x26, 0xCA},
			&z.H,
			0xCA,
		},
		{
			"LD L, n",
			[]byte{0x2E, 0xFE},
			&z.L,
			0xFE,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.LoadProgram(tc.program, 0)
			z.step()
			if *tc.register != tc.expected {
				t.Errorf("expected %x, got %x", tc.expected, *tc.register)
			}
		})
	}
}

func TestLDA(t *testing.T) {
	tbl := []testcase{
		{
			name:     "LD A,A",
			program:  []byte{0x7F},
			input:    Z80{A: 0xAA},
			expected: Z80{A: 0xAA},
		},
		{
			name:     "LD A,B",
			program:  []byte{0x78},
			input:    Z80{B: 0xDE},
			expected: Z80{A: 0xDE},
		},
		{
			name:     "LD A,C",
			program:  []byte{0x79},
			input:    Z80{C: 0xAD},
			expected: Z80{A: 0xAD},
		},
		{
			name:     "LD A,D",
			program:  []byte{0x7A},
			input:    Z80{D: 0xBE},
			expected: Z80{A: 0xBE},
		},
		{
			name:     "LD A,E",
			program:  []byte{0x7B},
			input:    Z80{E: 0xEF},
			expected: Z80{A: 0xEF},
		},
		{
			name:     "LD A,H",
			program:  []byte{0x7C},
			input:    Z80{H: 0xCA},
			expected: Z80{A: 0xCA},
		},
		{
			name:     "LD A,L",
			program:  []byte{0x7D},
			input:    Z80{L: 0xFE},
			expected: Z80{A: 0xFE},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			err := tc.input.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if tc.input.A != tc.expected.A {
				t.Errorf("expected %x, got %x", tc.expected.A, tc.input.A)
			}

			if tc.input.F != tc.expected.F {
				t.Errorf("expected flags %b, got %b", tc.expected.F, tc.input.F)
			}
		})
	}
}

func TestLDH(t *testing.T) {
	tbl := []testcase{
		{
			name:     "LD H, (HL)",
			program:  []byte{0x66, 0x00, 0x00, 0xCA, 0xFE},
			input:    Z80{L: 4},
			expected: Z80{H: 0xFE},
		},
		{
			name:     "LD H, (HL)",
			program:  []byte{0x66, 0x00, 0x00, 0xCA, 0xFE},
			input:    Z80{H: 4, L: 1},
			expected: Z80{H: 0x00},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			err := tc.input.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if tc.input.H != tc.expected.H {
				t.Errorf("expected %x, got %x", tc.expected.H, tc.input.H)
			}

			if tc.input.F != tc.expected.F {
				t.Errorf("expected flags %b, got %b", tc.expected.F, tc.input.F)
			}
		})
	}

}

// 8-Bit ALU
func TestADD(t *testing.T) {
	tbl := []testcase{
		{
			name:     "ADD A,A",
			program:  []byte{0x87},
			input:    Z80{A: 0x01},
			expected: Z80{A: 0x02, F: 0b00000000},
		},
		{
			name:     "ADD A,A ZFlag Set",
			program:  []byte{0x87},
			input:    Z80{A: 0x00},
			expected: Z80{A: 0x00, F: 0b10000000},
		},
		{
			name:     "ADD A,A HFlag Set",
			program:  []byte{0x87},
			input:    Z80{A: 0b00001000},
			expected: Z80{A: 0b00010000, F: 0b00100000},
		},
		{
			name:     "ADD A,A CFlag Set",
			program:  []byte{0x87},
			input:    Z80{A: 0b10000000},
			expected: Z80{A: 0b00000000, F: 0b00010000},
		},
		{
			name:     "ADD A,B",
			program:  []byte{0x80},
			input:    Z80{A: 0b00001100, B: 0b00001000},
			expected: Z80{A: 0b00010100, F: 0b00100000},
		},
		{
			name:     "ADD A,C",
			program:  []byte{0x81},
			input:    Z80{A: 0b00001100, C: 0b00010000},
			expected: Z80{A: 0b00011100, F: 0b00000000},
		},
		{
			name:     "ADD A,D",
			program:  []byte{0x82},
			input:    Z80{A: 0b10001100, D: 0b10001000},
			expected: Z80{A: 0b00010100, F: 0b00110000},
		},
		{
			name:     "ADD A,E",
			program:  []byte{0x83},
			input:    Z80{A: 0b00001100, E: 0b00001100},
			expected: Z80{A: 0b00011000, F: 0b00100000},
		},
		{
			name:     "ADD A,H",
			program:  []byte{0x84},
			input:    Z80{A: 0b00001100, H: 0b00000011},
			expected: Z80{A: 0b00001111, F: 0b00000000},
		},
		{
			name:     "ADD A,L",
			program:  []byte{0x85},
			input:    Z80{A: 0b00001100, L: 0b00001000},
			expected: Z80{A: 0b00010100, F: 0b00100000},
		},
		{
			name:     "ADD A,#",
			program:  []byte{0xC6, 0x10},
			input:    Z80{A: 0x20, F: 0b00010000},
			expected: Z80{A: 0x30, F: 0b00000000},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			err := tc.input.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if tc.input.A != tc.expected.A {
				t.Errorf("expected %x, got %x", tc.expected.A, tc.input.A)
			}

			if tc.input.F != tc.expected.F {
				t.Errorf("expected flags %b, got %b", tc.expected.F, tc.input.F)
			}
		})
	}
}

func TestADCA(t *testing.T) {
	tbl := []testcase{
		{
			name:     "ADC A,A",
			program:  []byte{0x8F},
			input:    Z80{A: 0x01},
			expected: Z80{A: 0x02, F: 0b00000000},
		},
		{
			name:     "ADC A,A CFlag Set",
			program:  []byte{0x8F},
			input:    Z80{A: 0x00, F: 0b00010000},
			expected: Z80{A: 0x01, F: 0b00000000},
		},
		{
			name:     "ADC A,B",
			program:  []byte{0x88},
			input:    Z80{A: 0b00001100, B: 0b00001000},
			expected: Z80{A: 0b00010100, F: 0b00100000},
		},
		{
			name:     "ADC A,C",
			program:  []byte{0x89},
			input:    Z80{A: 0b00001100, C: 0b00010000},
			expected: Z80{A: 0b00011100, F: 0b00000000},
		},
		{
			name:     "ADC A,D",
			program:  []byte{0x8A},
			input:    Z80{A: 0b10001100, D: 0b10001000},
			expected: Z80{A: 0b00010100, F: 0b00110000},
		},
		{
			name:     "ADC A,E",
			program:  []byte{0x8B},
			input:    Z80{A: 0b00001100, E: 0b00001100},
			expected: Z80{A: 0b00011000, F: 0b00100000},
		},
		{
			name:     "ADC A,H",
			program:  []byte{0x8C},
			input:    Z80{A: 0b00001100, H: 0b00000011},
			expected: Z80{A: 0b00001111, F: 0b00000000},
		},
		{
			name:     "ADC A,L",
			program:  []byte{0x8D},
			input:    Z80{A: 0b00001100, L: 0b00001000},
			expected: Z80{A: 0b00010100, F: 0b00100000},
		},
		{
			name:     "ADC A,#",
			program:  []byte{0xCE, 0x10},
			input:    Z80{A: 0x20, F: 0b00010000},
			expected: Z80{A: 0x31, F: 0b00000000},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			err := tc.input.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if tc.input.A != tc.expected.A {
				t.Errorf("expected %x, got %x", tc.expected.A, tc.input.A)
			}

			if tc.input.F != tc.expected.F {
				t.Errorf("expected flags %b, got %b", tc.expected.F, tc.input.F)
			}
		})
	}
}

// Stack Operations
func TestPUSH(t *testing.T) {
	tbl := []testcase{
		{
			name:       "PUSH AF",
			program:    []byte{0xF5},
			input:      Z80{A: 0xCA, F: 0xC0, SP: 0x0100},
			expected:   Z80{SP: 0x00FE, F: 0xC0},
			expected16: 0xCAC0,
		},
		{
			name:       "PUSH BC",
			program:    []byte{0xC5},
			input:      Z80{B: 0xDE, C: 0xAD, SP: 0x0100},
			expected:   Z80{SP: 0x00FE},
			expected16: 0xDEAD,
		},
		{
			name:       "PUSH DE",
			program:    []byte{0xD5},
			input:      Z80{D: 0xBE, E: 0xEF, SP: 0x0100},
			expected:   Z80{SP: 0x00FE},
			expected16: 0xBEEF,
		},
		{
			name:       "PUSH HL",
			program:    []byte{0xE5},
			input:      Z80{H: 0xF0, L: 0x0D, SP: 0x0100},
			expected:   Z80{SP: 0x00FE},
			expected16: 0xF00D,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			err := tc.input.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if tc.input.SP != tc.expected.SP {
				t.Errorf("expected SP %x, got %x", tc.expected.SP, tc.input.SP)
			}

			hi := m.Read(tc.input.SP)
			lo := m.Read(tc.input.SP + 1)

			if value := pair(hi, lo); value != tc.expected16 {
				t.Errorf("expected value %x, got %x", tc.expected16, value)
			}

			if tc.input.F != tc.expected.F {
				t.Errorf("expected flags %b, got %b", tc.expected.F, tc.input.F)
			}
		})
	}
}

func TestPOP(t *testing.T) {
	tbl := []testcase{
		{
			name:     "POP AF",
			program:  []byte{0xF1},
			input16:  0xCAC0,
			input:    Z80{SP: 0x0100},
			expected: Z80{A: 0xCA, F: 0xC0, SP: 0x0102},
		},
		{
			name:     "POP BC",
			program:  []byte{0xC1},
			input16:  0xDEAD,
			input:    Z80{SP: 0x0100},
			expected: Z80{B: 0xDE, C: 0xAD, SP: 0x0102},
		},
		{
			name:     "POP DE",
			program:  []byte{0xD1},
			input16:  0xBEEF,
			input:    Z80{SP: 0x0100},
			expected: Z80{D: 0xBE, E: 0xEF, SP: 0x0102},
		},
		{
			name:     "POP HL",
			program:  []byte{0xE1},
			input16:  0xF00D,
			input:    Z80{SP: 0x0100},
			expected: Z80{H: 0xF0, L: 0x0D, SP: 0x0102},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			hi := byte(tc.input16 / 0x100)
			lo := byte(tc.input16 % 0x100)
			m.Write(tc.input.SP, hi)
			m.Write(tc.input.SP+1, lo)

			err := tc.input.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if tc.input.SP != tc.expected.SP {
				t.Errorf("expected SP %x, got %x", tc.expected.SP, tc.input.SP)
			}

			if tc.input.AF() != tc.expected.AF() {
				t.Errorf("expected AF %x, got %x", tc.expected.AF(), tc.input.AF())
			}

			if tc.input.BC() != tc.expected.BC() {
				t.Errorf("expected BC %x, got %x", tc.expected.BC(), tc.input.BC())
			}

			if tc.input.DE() != tc.expected.DE() {
				t.Errorf("expected DE %x, got %x", tc.expected.DE(), tc.input.DE())
			}

			if tc.input.HL() != tc.expected.HL() {
				t.Errorf("expected HL %x, got %x", tc.expected.HL(), tc.input.HL())
			}

			if tc.input.F != tc.expected.F {
				t.Errorf("expected flags %b, got %b", tc.expected.F, tc.input.F)
			}
		})
	}
}
