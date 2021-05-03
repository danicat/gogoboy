package main

import "testing"

func TestFetch(t *testing.T) {
	input := []byte{0xDE, 0xAD}
	expected := byte(0xDE)
	z := NewZ80(input)
	op := z.fetch()
	if op != expected {
		t.Fatalf("expected %x, got %x", expected, op)
	}
	if z.PC != 1 {
		t.Fatalf("expected %x, got %x", 1, z.PC)
	}
}

func TestNOP(t *testing.T) {
	input := []byte{0, 0}
	expected := int16(2)
	z := NewZ80(input)
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
	z := NewZ80(input)
	z.step()
	if z.PC != 2 {
		t.Fatalf("expected PC=%d, got PC=%d", 2, z.PC)
	}
	if z.B != expected {
		t.Fatalf("expected B=%d, got %d", expected, z.B)
	}
}

func TestLDn(t *testing.T) {
	z := NewZ80(nil)
	tbl := []struct {
		name     string
		program  []byte
		register *byte
		expected byte
	}{
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

	for _, testcase := range tbl {
		t.Run(testcase.name, func(t *testing.T) {
			z.Reset()
			z.LoadProgram(testcase.program)
			z.step()
			if *testcase.register != testcase.expected {
				t.Errorf("expected %x, got %x", testcase.expected, *testcase.register)
			}
		})
	}
}

func TestLDA(t *testing.T) {
	z := NewZ80(nil)

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

	for _, testcase := range tbl {
		t.Run(testcase.name, func(t *testing.T) {
			z.Reset()
			z.A = 0xAA
			z.B = 0xDE
			z.C = 0xAD
			z.D = 0xBE
			z.E = 0xEF
			z.H = 0xCA
			z.L = 0xFE

			z.LoadProgram(testcase.program)

			z.step()
			if z.A != testcase.expected {
				t.Errorf("expected %x, got %x", testcase.expected, z.A)
			}
		})
	}
}

// 8-Bit ALU
func TestADD(t *testing.T) {
	z := NewZ80(nil)

	tbl := []struct {
		name      string
		program   []byte
		A         byte
		B         byte
		C         byte
		D         byte
		E         byte
		H         byte
		L         byte
		expectedA byte
		expectedF byte
	}{
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
	}

	for _, testcase := range tbl {
		t.Run(testcase.name, func(t *testing.T) {
			z.Reset()
			z.A = testcase.A
			z.B = testcase.B
			z.C = testcase.C
			z.D = testcase.D
			z.E = testcase.E
			z.H = testcase.H
			z.L = testcase.L

			z.LoadProgram(testcase.program)

			z.step()

			if z.A != testcase.expectedA {
				t.Errorf("expected %x, got %x", testcase.expectedA, z.A)
			}

			if z.F != testcase.expectedF {
				t.Errorf("expected flags %b, got %b", testcase.expectedF, z.F)
			}

		})
	}
}
