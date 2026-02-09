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
	headerLength     = len(magicBytes) + 1 + saltLength + wrappedDEKLength // magic + version + salt + edek
)

func PackagePayload(payload EncryptedDataPayload) []byte {
	header := make([]byte, headerLength)

	// Add magic bytes
	copy(header, []byte(magicBytes))
	// Add version
	header[len(magicBytes)] = formatVersion
	// Add salt and DEK
	copy(header[len(magicBytes)+1:], payload.Salt)
	copy(header[len(magicBytes)+1+len(payload.Salt):], payload.DEK)

	buffer := make([]byte, headerLength+len(payload.Payload))
	copy(buffer, header)
	copy(buffer[len(header):], payload.Payload)

	return buffer
}

func UnpackagePayload(data []byte) (EncryptedDataPayload, error) {
	if len(data) < headerLength {
		return EncryptedDataPayload{}, errors.New("data too short")
	}

	// Verify magic bytes
	if !bytes.Equal(data[:len(magicBytes)], []byte(magicBytes)) {
		return EncryptedDataPayload{}, errors.New("invalid file format: magic bytes mismatch")
	}

	// Verify version
	if data[len(magicBytes)] != formatVersion {
		return EncryptedDataPayload{}, errors.New("unsupported format version")
	}

	offset := len(magicBytes) + 1
	salt := data[offset : offset+saltLength]
	edek := data[offset+saltLength : offset+saltLength+wrappedDEKLength]
	payload := data[headerLength:]

	return EncryptedDataPayload{
		Salt:    Salt(salt),
		DEK:     WrappedDEK(edek),
		Payload: Ciphertext(payload),
	}, nil
}
