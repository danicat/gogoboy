package main

import "testing"

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

// 8-Bit Load
func TestLDB(t *testing.T) {
	input := []byte{0x06, 0xDE}
	var expected byte = 0xDE
	z := NewZ80()
	z.LoadProgram(input, 0)
	z.step()
	if z.PC != 2 {
		t.Fatalf("expected PC=%d, got PC=%d", 2, z.PC)
	}
	if z.B != expected {
		t.Fatalf("expected B=%d, got %d", expected, z.B)
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
	z := NewZ80()

	tbl := []testcase{
		{
			name:      "LD A,A",
			program:   []byte{0x7F},
			A:         0xAA,
			expectedA: 0xAA,
		},
		{
			name:      "LD A,B",
			program:   []byte{0x78},
			B:         0xDE,
			expectedA: 0xDE,
		},
		{
			name:      "LD A,C",
			program:   []byte{0x79},
			C:         0xAD,
			expectedA: 0xAD,
		},
		{
			name:      "LD A,D",
			program:   []byte{0x7A},
			D:         0xBE,
			expectedA: 0xBE,
		},
		{
			name:      "LD A,E",
			program:   []byte{0x7B},
			E:         0xEF,
			expectedA: 0xEF,
		},
		{
			name:      "LD A,H",
			program:   []byte{0x7C},
			H:         0xCA,
			expectedA: 0xCA,
		},
		{
			name:      "LD A,L",
			program:   []byte{0x7D},
			L:         0xFE,
			expectedA: 0xFE,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.A = tc.A
			z.F = tc.F
			z.B = tc.B
			z.C = tc.C
			z.D = tc.D
			z.E = tc.E
			z.H = tc.H
			z.L = tc.L

			z.LoadProgram(tc.program, 0)

			err := z.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if z.A != tc.expectedA {
				t.Errorf("expected %x, got %x", tc.expectedA, z.A)
			}

			if z.F != tc.expectedF {
				t.Errorf("expected flags %b, got %b", tc.expectedF, z.F)
			}
		})
	}
}

func TestLDH(t *testing.T) {
	z := NewZ80()

	tbl := []testcase{
		{
			name:      "LD H, (HL)",
			program:   []byte{0x66, 0x00, 0x00, 0xCA, 0xFE},
			L:         4,
			expectedH: 0xFE,
		},
		{
			name:      "LD H, (HL)",
			program:   []byte{0x66, 0x00, 0x00, 0xCA, 0xFE},
			H:         4,
			L:         1,
			expectedH: 0x00,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.A = tc.A
			z.F = tc.F
			z.B = tc.B
			z.C = tc.C
			z.D = tc.D
			z.E = tc.E
			z.H = tc.H
			z.L = tc.L

			z.LoadProgram(tc.program, 0)

			err := z.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if z.H != tc.expectedH {
				t.Errorf("expected %x, got %x", tc.expectedH, z.H)
			}

			if z.F != tc.expectedF {
				t.Errorf("expected flags %b, got %b", tc.expectedF, z.F)
			}
		})
	}

}

// 8-Bit ALU
func TestADD(t *testing.T) {
	z := NewZ80()

	tbl := []testcase{
		{
			name:      "ADD A,A",
			program:   []byte{0x87},
			A:         0x01,
			expectedA: 0x02,
			expectedF: 0b00000000,
		},
		{
			name:      "ADD A,A ZFlag Set",
			program:   []byte{0x87},
			A:         0x00,
			expectedA: 0x00,
			expectedF: 0b10000000,
		},
		{
			name:      "ADD A,A HFlag Set",
			program:   []byte{0x87},
			A:         0b00001000,
			expectedA: 0b00010000,
			expectedF: 0b00100000,
		},
		{
			name:      "ADD A,A CFlag Set",
			program:   []byte{0x87},
			A:         0b10000000,
			expectedA: 0b00000000,
			expectedF: 0b00010000,
		},
		{
			name:      "ADD A,B",
			program:   []byte{0x80},
			A:         0b00001100,
			B:         0b00001000,
			expectedA: 0b00010100,
			expectedF: 0b00100000,
		},
		{
			name:      "ADD A,C",
			program:   []byte{0x81},
			A:         0b00001100,
			C:         0b00010000,
			expectedA: 0b00011100,
			expectedF: 0b00000000,
		},
		{
			name:      "ADD A,D",
			program:   []byte{0x82},
			A:         0b10001100,
			D:         0b10001000,
			expectedA: 0b00010100,
			expectedF: 0b00110000,
		},
		{
			name:      "ADD A,E",
			program:   []byte{0x83},
			A:         0b00001100,
			E:         0b00001100,
			expectedA: 0b00011000,
			expectedF: 0b00100000,
		},
		{
			name:      "ADD A,H",
			program:   []byte{0x84},
			A:         0b00001100,
			H:         0b00000011,
			expectedA: 0b00001111,
			expectedF: 0b00000000,
		},
		{
			name:      "ADD A,L",
			program:   []byte{0x85},
			A:         0b00001100,
			L:         0b00001000,
			expectedA: 0b00010100,
			expectedF: 0b00100000,
		},
		{
			name:      "ADD A,#",
			program:   []byte{0xC6, 0x10},
			A:         0x20,
			F:         0b00010000,
			expectedA: 0x30,
			expectedF: 0b00000000,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.A = tc.A
			z.B = tc.B
			z.C = tc.C
			z.D = tc.D
			z.E = tc.E
			z.H = tc.H
			z.L = tc.L

			z.LoadProgram(tc.program, 0)

			err := z.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if z.A != tc.expectedA {
				t.Errorf("expected %x, got %x", tc.expectedA, z.A)
			}

			if z.F != tc.expectedF {
				t.Errorf("expected flags %b, got %b", tc.expectedF, z.F)
			}
		})
	}
}

func TestADCA(t *testing.T) {
	z := NewZ80()

	tbl := []testcase{
		{
			name:      "ADC A,A",
			program:   []byte{0x8F},
			A:         0x01,
			expectedA: 0x02,
			expectedF: 0b00000000,
		},
		{
			name:      "ADC A,A CFlag Set",
			program:   []byte{0x8F},
			A:         0x00,
			F:         0b00010000,
			expectedA: 0x01,
			expectedF: 0b00000000,
		},
		{
			name:      "ADC A,B",
			program:   []byte{0x88},
			A:         0b00001100,
			B:         0b00001000,
			expectedA: 0b00010100,
			expectedF: 0b00100000,
		},
		{
			name:      "ADC A,C",
			program:   []byte{0x89},
			A:         0b00001100,
			C:         0b00010000,
			expectedA: 0b00011100,
			expectedF: 0b00000000,
		},
		{
			name:      "ADC A,D",
			program:   []byte{0x8A},
			A:         0b10001100,
			D:         0b10001000,
			expectedA: 0b00010100,
			expectedF: 0b00110000,
		},
		{
			name:      "ADC A,E",
			program:   []byte{0x8B},
			A:         0b00001100,
			E:         0b00001100,
			expectedA: 0b00011000,
			expectedF: 0b00100000,
		},
		{
			name:      "ADC A,H",
			program:   []byte{0x8C},
			A:         0b00001100,
			H:         0b00000011,
			expectedA: 0b00001111,
			expectedF: 0b00000000,
		},
		{
			name:      "ADC A,L",
			program:   []byte{0x8D},
			A:         0b00001100,
			L:         0b00001000,
			expectedA: 0b00010100,
			expectedF: 0b00100000,
		},
		{
			name:      "ADC A,#",
			program:   []byte{0xCE, 0x10},
			A:         0x20,
			F:         0b00010000,
			expectedA: 0x31,
			expectedF: 0b00000000,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.A = tc.A
			z.F = tc.F
			z.B = tc.B
			z.C = tc.C
			z.D = tc.D
			z.E = tc.E
			z.H = tc.H
			z.L = tc.L

			z.LoadProgram(tc.program, 0)

			err := z.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if z.A != tc.expectedA {
				t.Errorf("expected %x, got %x", tc.expectedA, z.A)
			}

			if z.F != tc.expectedF {
				t.Errorf("expected flags %b, got %b", tc.expectedF, z.F)
			}
		})
	}
}

// Stack Operations
func TestPUSH(t *testing.T) {
	z := NewZ80()

	tbl := []testcase{
		{
			name:       "PUSH AF",
			program:    []byte{0xF5},
			A:          0xCA,
			F:          0xC0,
			SP:         0x0100,
			expectedSP: 0x00FE,
			expectedF:  0xC0,
			expected16: 0xCAC0,
		},
		{
			name:       "PUSH BC",
			program:    []byte{0xC5},
			B:          0xDE,
			C:          0xAD,
			SP:         0x0100,
			expectedSP: 0x00FE,
			expected16: 0xDEAD,
		},
		{
			name:       "PUSH DE",
			program:    []byte{0xD5},
			D:          0xBE,
			E:          0xEF,
			SP:         0x0100,
			expectedSP: 0x00FE,
			expected16: 0xBEEF,
		},
		{
			name:       "PUSH HL",
			program:    []byte{0xE5},
			H:          0xF0,
			L:          0x0D,
			SP:         0x0100,
			expectedSP: 0x00FE,
			expected16: 0xF00D,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.SP = tc.SP
			z.A = tc.A
			z.F = tc.F
			z.B = tc.B
			z.C = tc.C
			z.D = tc.D
			z.E = tc.E
			z.H = tc.H
			z.L = tc.L

			z.LoadProgram(tc.program, 0)

			err := z.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if z.SP != tc.expectedSP {
				t.Errorf("expected SP %x, got %x", tc.expectedSP, z.SP)
			}

			hi := z.ram.Read(z.SP)
			lo := z.ram.Read(z.SP + 1)

			if value := pair(hi, lo); value != tc.expected16 {
				t.Errorf("expected value %x, got %x", tc.expected16, value)
			}

			if z.F != tc.expectedF {
				t.Errorf("expected flags %b, got %b", tc.expectedF, z.F)
			}
		})
	}
}

func TestPOP(t *testing.T) {
	z := NewZ80()

	tbl := []testcase{
		{
			name:       "POP AF",
			program:    []byte{0xF1},
			input16:    0xCAC0,
			SP:         0x0100,
			expectedA:  0xCA,
			expectedF:  0xC0,
			expectedSP: 0x0102,
		},
		{
			name:       "POP BC",
			program:    []byte{0xC1},
			input16:    0xDEAD,
			SP:         0x0100,
			expectedB:  0xDE,
			expectedC:  0xAD,
			expectedSP: 0x0102,
		},
		{
			name:       "POP DE",
			program:    []byte{0xD1},
			input16:    0xBEEF,
			SP:         0x0100,
			expectedD:  0xBE,
			expectedE:  0xEF,
			expectedSP: 0x0102,
		},
		{
			name:       "POP HL",
			program:    []byte{0xE1},
			input16:    0xF00D,
			SP:         0x0100,
			expectedH:  0xF0,
			expectedL:  0x0D,
			expectedSP: 0x0102,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.SP = tc.SP
			z.A = tc.A
			z.F = tc.F
			z.B = tc.B
			z.C = tc.C
			z.D = tc.D
			z.E = tc.E
			z.H = tc.H
			z.L = tc.L

			z.LoadProgram(tc.program, 0)

			hi := byte(tc.input16 / 0x100)
			lo := byte(tc.input16 % 0x100)
			z.ram.Write(z.SP, hi)
			z.ram.Write(z.SP+1, lo)

			err := z.step()
			if err != nil {
				t.Errorf("expected no error, got: %s", err)
			}

			if z.SP != tc.expectedSP {
				t.Errorf("expected SP %x, got %x", tc.expectedSP, z.SP)
			}

			if eAF := pair(tc.expectedA, tc.expectedF); z.GetAF() != eAF {
				t.Errorf("expected AF %x, got %x", eAF, z.GetAF())
			}

			if eBC := pair(tc.expectedB, tc.expectedC); z.GetBC() != eBC {
				t.Errorf("expected BC %x, got %x", eBC, z.GetBC())
			}

			if eDE := pair(tc.expectedD, tc.expectedE); z.GetDE() != eDE {
				t.Errorf("expected DE %x, got %x", eDE, z.GetDE())
			}

			if eHL := pair(tc.expectedH, tc.expectedL); z.GetHL() != eHL {
				t.Errorf("expected HL %x, got %x", eHL, z.GetHL())
			}

			if z.F != tc.expectedF {
				t.Errorf("expected flags %b, got %b", tc.expectedF, z.F)
			}
		})
	}
}
