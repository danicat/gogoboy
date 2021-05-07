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
