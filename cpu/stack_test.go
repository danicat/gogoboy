package cpu

import (
	"testing"

	"github.com/danicat/gogoboy/memory"
)

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
