package ultrahash

import (
	"bytes"
	"testing"
)

func TestUltraHash(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		salt  []byte
	}{
		{"nil", nil, nil},
		{"empty", []byte(""), []byte("")},
		{"password", []byte("password"), nil},
		{"salted password", []byte("password"), []byte("pepper")},
		{"1 MiB", bytes.Repeat([]byte{0xff}, 1<<20), bytes.Repeat([]byte{0x00}, 1<<20)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UltraHash(tt.input, tt.salt); got != 42 {
				t.Errorf("UltraHash() = %d, want 42", got)
			}
		})
	}
}
