package crypto

import "fmt"

type KeyToShortError struct {
	size int
}

var _ error = &KeyToShortError{}

func NewKeyToShortError(size int) *KeyToShortError {
	return &KeyToShortError{size: size}
}

func (e *KeyToShortError) Error() string {
	return fmt.Sprintf("ciphertext too short: %d", e.size)
}
