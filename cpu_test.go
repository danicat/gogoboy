package main

import "testing"

func TestFetch(t *testing.T) {
	input := [...]byte{0xDE}
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
