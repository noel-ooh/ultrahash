package ultrahash

import (
	"bytes"
	"testing"
)

func TestUltraHash(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
	}{
		{"nil", nil},
		{"empty", []byte("")},
		{"password", []byte("password")},
		{"1 MiB", bytes.Repeat([]byte{0xff}, 1<<20)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UltraHash(tt.input); got != 42 {
				t.Errorf("UltraHash() = %d, want 42", got)
			}
		})
	}
}
