package memory_test

import (
	"testing"

	"github.com/danicat/gogoboy/memory"
)

func TestLoadProgram(t *testing.T) {
	program := []byte{0, 1, 2, 3}
	m := memory.NewMemory()
	m.LoadProgram(program, 0)
	b := m.Read(3)
	if b != 3 {
		t.Fatalf("expected 3, got %d", b)
	}
}
