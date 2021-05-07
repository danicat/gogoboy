package cpu

import (
	"testing"

	"github.com/danicat/gogoboy/memory"
)

func TestCALLcc(t *testing.T) {
	tbl := []testcase{
		{
			name:     "CALL Z, nn - Z",
			program:  []byte{0x87, 0xCC, 0x01, 0x00},
			input:    Z80{SP: 0xFFFE},
			expected: Z80{F: 0b10000000, PC: 0x0100, SP: 0xFFFC},
		},
		{
			name:     "CALL Z, nn - NZ",
			program:  []byte{0x87, 0xCC, 0x01, 0x00},
			input:    Z80{A: 0x01, SP: 0xFFFE},
			expected: Z80{A: 0x02, F: 0b00000000, PC: 0x0004, SP: 0xFFFE},
		},
		{
			name:     "CALL NZ, nn - Z",
			program:  []byte{0x87, 0xC4, 0x01, 0x00},
			input:    Z80{SP: 0xFFFE},
			expected: Z80{F: 0b10000000, PC: 0x0004, SP: 0xFFFE},
		},
		{
			name:     "CALL NZ, nn - NZ",
			program:  []byte{0x87, 0xC4, 0x01, 0x00},
			input:    Z80{A: 0x01, SP: 0xFFFE},
			expected: Z80{A: 0x02, F: 0b00000000, PC: 0x0100, SP: 0xFFFC},
		},
		{
			name:     "CALL C, nn - C",
			program:  []byte{0x80, 0xDC, 0x01, 0x00},
			input:    Z80{A: 0xFF, B: 0x01, SP: 0xFFFE},
			expected: Z80{B: 0x01, F: 0b00110000, PC: 0x0100, SP: 0xFFFC},
		},
		{
			name:     "CALL C, nn - NC",
			program:  []byte{0x80, 0xDC, 0x01, 0x00},
			input:    Z80{A: 0x01, SP: 0xFFFE},
			expected: Z80{A: 0x01, F: 0b00000000, PC: 0x0004, SP: 0xFFFE},
		},
		{
			name:     "CALL NC, nn - C",
			program:  []byte{0x80, 0xD4, 0x01, 0x00},
			input:    Z80{A: 0xFF, B: 0x01, SP: 0xFFFE},
			expected: Z80{B: 0x01, F: 0b00110000, PC: 0x0004, SP: 0xFFFE},
		},
		{
			name:     "CALL NC, nn - NC",
			program:  []byte{0x80, 0xD4, 0x01, 0x00},
			input:    Z80{A: 0x01, SP: 0xFFFE},
			expected: Z80{A: 0x01, F: 0b00000000, PC: 0x0100, SP: 0xFFFC},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			m := memory.NewMemory()
			tc.input.ram = m
			tc.input.LoadProgram(tc.program, 0)

			tc.input.step()
			tc.input.step()

			if tc.input.SP != tc.expected.SP {
				t.Errorf("expected SP %x, got %x", tc.expected.SP, tc.input.SP)
			}

			if tc.input.PC != tc.expected.PC {
				t.Errorf("expected PC %x, got %x", tc.expected.PC, tc.input.PC)
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
