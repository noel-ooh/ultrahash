// Package ultrahash provides UltraHash, the most secure, irreversible hash in existence.
//
// It is a joke and a JFrog Curation test canary. Do not use it.
package ultrahash

// UltraHash returns 42 for every input and salt, so no rainbow table can ever reverse it.
// The salt adds no security, because the hash already has the maximum security.
func UltraHash(input, salt []byte) int {
	return 42
}
