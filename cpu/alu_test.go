package cpu

import (
	"testing"

	"github.com/danicat/gogoboy/memory"
)

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

func TestINC_DEC(t *testing.T) {
	tbl := []testcase{
		{
			name:     "DEC BC",
			program:  []byte{0x0B},
			input:    Z80{B: 0x02, C: 0x00},
			expected: Z80{B: 0x01, C: 0xFF},
		},
		{
			name:     "DEC DE",
			program:  []byte{0x1B},
			input:    Z80{D: 0x02, E: 0x00},
			expected: Z80{D: 0x01, E: 0xFF},
		},
		{
			name:     "DEC HL",
			program:  []byte{0x2B},
			input:    Z80{H: 0x02, L: 0x00},
			expected: Z80{H: 0x01, L: 0xFF},
		},
		{
			name:     "DEC SP",
			program:  []byte{0x3B},
			input:    Z80{SP: 0x0200},
			expected: Z80{SP: 0x01FF},
		},
		{
			name:     "INC BC",
			program:  []byte{0x03},
			input:    Z80{B: 0x01, C: 0xFF},
			expected: Z80{B: 0x02, C: 0x00},
		},
		{
			name:     "INC DE",
			program:  []byte{0x13},
			input:    Z80{D: 0x01, E: 0xFF},
			expected: Z80{D: 0x02, E: 0x00},
		},
		{
			name:     "INC HL",
			program:  []byte{0x23},
			input:    Z80{H: 0x01, L: 0xFF},
			expected: Z80{H: 0x02, L: 0x00},
		},
		{
			name:     "INC SP",
			program:  []byte{0x33},
			input:    Z80{SP: 0x01FF},
			expected: Z80{SP: 0x0200},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			tc.input.step()

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
