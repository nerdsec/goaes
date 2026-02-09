package internal

import "errors"

type (
	KEK        []byte
	DEK        []byte
	WrappedDEK []byte
	Ciphertext []byte
	Salt       []byte
)

type EncryptedDataPayload struct {
	DEK     WrappedDEK
	Salt    Salt
	Payload Ciphertext
}

var (
	aadWrapDEK  = []byte("wrap:dek:v1")
	aadDataMsg  = []byte("data:msg:v1")
	errBadKeyLn = errors.New("invalid key length: must be 16, 24, or 32 bytes")
)

func clear(b []byte) {
	for i := range b {
		b[i] = 0
	}
}
