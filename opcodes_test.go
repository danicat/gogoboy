package main

import "testing"

func TestNOP(t *testing.T) {
	input := []byte{0, 0}
	expected := int16(2)
	z := NewZ80()
	z.LoadProgram(input)
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
	z.LoadProgram(input)
	z.step()
	if z.PC != 2 {
		t.Fatalf("expected PC=%d, got PC=%d", 2, z.PC)
	}
	if z.B != expected {
		t.Fatalf("expected B=%d, got %d", expected, z.B)
	}
}

func TestLDn(t *testing.T) {
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
			"LD B,n",
			[]byte{0x06, 0xDE},
			&z.B,
			0xDE,
		},
		{
			"LD C,n",
			[]byte{0x0E, 0xAD},
			&z.C,
			0xAD,
		},
		{
			"LD D,n",
			[]byte{0x16, 0xBE},
			&z.D,
			0xBE,
		},
		{
			"LD E,n",
			[]byte{0x1E, 0xEF},
			&z.E,
			0xEF,
		},
		{
			"LD H,n",
			[]byte{0x26, 0xCA},
			&z.H,
			0xCA,
		},
		{
			"LD L,n",
			[]byte{0x2E, 0xFE},
			&z.L,
			0xFE,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.LoadProgram(tc.program)
			z.step()
			if *tc.register != tc.expected {
				t.Errorf("expected %x, got %x", tc.expected, *tc.register)
			}
		})
	}
}

func TestLDA(t *testing.T) {
	z := NewZ80()

	tbl := []struct {
		name     string
		program  []byte
		expected byte
	}{
		{
			"LD A,A",
			[]byte{0x7F},
			0xAA,
		},
		{
			"LD A,B",
			[]byte{0x78},
			0xDE,
		},
		{
			"LD A,C",
			[]byte{0x79},
			0xAD,
		},
		{
			"LD A,D",
			[]byte{0x7A},
			0xBE,
		},
		{
			"LD A,E",
			[]byte{0x7B},
			0xEF,
		},
		{
			"LD A,H",
			[]byte{0x7C},
			0xCA,
		},
		{
			"LD A,L",
			[]byte{0x7D},
			0xFE,
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {
			z.Reset()
			z.A = 0xAA
			z.B = 0xDE
			z.C = 0xAD
			z.D = 0xBE
			z.E = 0xEF
			z.H = 0xCA
			z.L = 0xFE

			z.LoadProgram(tc.program)

			z.step()
			if z.A != tc.expected {
				t.Errorf("expected %x, got %x", tc.expected, z.A)
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

			z.LoadProgram(tc.program)

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

			z.LoadProgram(tc.program)

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
