package cpu

import "testing"

type testcase struct {
	name       string
	program    []byte
	input      Z80
	expected   Z80
	input16    uint16
	expected16 uint16
}

func TestFetch(t *testing.T) {
	input := []byte{0xDE, 0xAD}
	expected := byte(0xDE)
	z := NewZ80()
	z.LoadProgram(input, 0)
	op := z.fetch()
	if op != expected {
		t.Fatalf("expected %x, got %x", expected, op)
	}
	if z.PC != 1 {
		t.Fatalf("expected %x, got %x", 1, z.PC)
	}
}

func TestCycleLimit(t *testing.T) {
	tbl := []testcase{
		{
			name:     "LD A,0x1F LD B,0x21 ADD A,B",
			program:  []byte{0x3E, 0x1F, 0x06, 0x21, 0x80},
			expected: Z80{A: 0x40, F: 0b00100000, cycles: 100},
		},
	}

	z := NewZ80()

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.SetMaxCycles(100)
			z.LoadProgram(tc.program, 0)
			z.Run()

			if z.A != tc.expected.A {
				t.Errorf("expected A=%x, got %x", tc.expected.A, z.A)
			}
			if z.F != tc.expected.F {
				t.Errorf("expected flags %b, got %b", tc.expected.F, z.F)
			}
			if z.cycles != tc.expected.cycles {
				t.Errorf("expected cycles=%d, got %d", tc.expected.cycles, z.cycles)
			}
		})
	}
}

func TestNintendoBootProgram(t *testing.T) {
	input := []byte{
		0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
		0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
		0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E,
	}

	z := NewZ80()
	z.LoadProgram(input, 0)
	z.SetMaxCycles(1000)
	err := z.Run()
	if err != nil {
		t.Errorf("expected no errors, got: %s", err)
	}
}
