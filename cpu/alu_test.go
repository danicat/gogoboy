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
