package internal

import (
	"bytes"
	"errors"
)

const (
	magicBytes       = "GOAES"
	formatVersion    = 0x01
	saltLength       = 32
	wrappedDEKLength = 60
)

func PackagePayload(payload EncryptedDataPayload) []byte {
	var buf bytes.Buffer

	magic := []byte(magicBytes)

	// Start marker
	buf.Write(magic)

	// Header contents
	buf.WriteByte(formatVersion)
	buf.Write(payload.Salt)
	buf.Write(payload.DEK)

	// End marker
	buf.Write(magic)

	// Ciphertext payload
	buf.Write(payload.Payload)

	return buf.Bytes()
}

func UnpackagePayload(data []byte) (EncryptedDataPayload, error) {
	magic := []byte(magicBytes)
	magicLength := len(magic)

	// Must at least contain: magic + version + salt + dek + magic
	minHeaderSize := magicLength + 1 + saltLength + wrappedDEKLength + magicLength
	if len(data) < minHeaderSize {
		return EncryptedDataPayload{}, errors.New("data too short")
	}

	// Verify starting magic
	if !bytes.Equal(data[:magicLength], magic) {
		return EncryptedDataPayload{}, errors.New("invalid file format: missing starting magic bytes")
	}

	offset := magicLength

	// Version
	version := data[offset]
	if version != formatVersion {
		return EncryptedDataPayload{}, errors.New("unsupported format version")
	}
	offset++

	// Salt
	if len(data) < offset+saltLength {
		return EncryptedDataPayload{}, errors.New("truncated salt")
	}
	salt := data[offset : offset+saltLength]
	offset += saltLength

	// Wrapped DEK
	if len(data) < offset+wrappedDEKLength {
		return EncryptedDataPayload{}, errors.New("truncated wrapped DEK")
	}
	edek := data[offset : offset+wrappedDEKLength]
	offset += wrappedDEKLength

	// Verify ending magic
	if len(data) < offset+magicLength || !bytes.Equal(data[offset:offset+magicLength], magic) {
		return EncryptedDataPayload{}, errors.New("invalid file format: missing ending magic bytes")
	}
	offset += magicLength

	// Remaining bytes are ciphertext payload
	payload := data[offset:]

	return EncryptedDataPayload{
		Salt:    Salt(salt),
		DEK:     WrappedDEK(edek),
		Payload: Ciphertext(payload),
	}, nil
}
