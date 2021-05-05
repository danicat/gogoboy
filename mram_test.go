package main

import "testing"

func TestLoadProgram(t *testing.T) {
	program := []byte{0, 1, 2, 3}
	m := NewMRAM()
	m.LoadProgram(program)
	b := m.Read(3)
	if b != 3 {
		t.Fatalf("expected 3, got %d", b)
	}
}
