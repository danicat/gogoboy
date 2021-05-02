package main

import "testing"

func TestFetch(t *testing.T) {
	input := []byte{0xDE}
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
		t.Fatalf("expected %d, got %d", 2, z.PC)
	}
}
