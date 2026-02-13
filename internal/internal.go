package internal

import (
	"errors"
	"fmt"
)

type (
	KEK        []byte
	DEK        []byte
	WrappedDEK []byte
	Ciphertext []byte
	Salt       []byte
)

type EncryptedDataPayload struct {
	WrappedDEK WrappedDEK
	Salt       Salt
	Payload    Ciphertext
}

var (
	aadWrapDEK  = []byte("wrap:dek:v1")
	aadDataMsg  = []byte("data:msg:v1")
	errBadKeyLn = errors.New("invalid key length: must be 16, 24, or 32 bytes")
)

const (
	saltLen       = 32
	wrappedDEKLen = 60
)

func Clear(b []byte) {
	clear(b)
}

func validPayloadLengths(salt Salt, edek WrappedDEK) error {
	if len(salt) != saltLen {
		return fmt.Errorf("unexpected salt length: %d", len(salt))
	}

	if len(edek) != wrappedDEKLen {
		return fmt.Errorf("unexpected wrapped DEK length: %d", len(edek))
	}

	return nil
}
