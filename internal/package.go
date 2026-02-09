package internal

import (
	"bytes"
	"errors"
)

const (
	magicBytes    = "GOAES"
	formatVersion = 0x01
	headerLength  = len(magicBytes) + 1 + saltLen + wrappedDEKLen // magic + version + salt + edek
)

func PackagePayload(payload EncryptedDataPayload) ([]byte, error) {
	if err := validPayloadLengths(payload.Salt, payload.WrappedDEK); err != nil {
		return nil, err
	}

	buf := make([]byte, headerLength+len(payload.Payload))
	copy(buf, []byte(magicBytes))
	buf[len(magicBytes)] = formatVersion
	copy(buf[len(magicBytes)+1:], payload.Salt)
	copy(buf[len(magicBytes)+1+saltLen:], payload.WrappedDEK)
	copy(buf[headerLength:], payload.Payload)

	return buf, nil
}

func UnpackagePayload(data []byte) (EncryptedDataPayload, error) {
	if len(data) < headerLength {
		return EncryptedDataPayload{}, errors.New("data too short")
	}

	if !bytes.Equal(data[:len(magicBytes)], []byte(magicBytes)) {
		return EncryptedDataPayload{}, errors.New("invalid file format: magic bytes mismatch")
	}

	if data[len(magicBytes)] != formatVersion {
		return EncryptedDataPayload{}, errors.New("unsupported format version")
	}

	offset := len(magicBytes) + 1

	salt := make([]byte, saltLen)
	copy(salt, data[offset:offset+saltLen])

	edek := make([]byte, wrappedDEKLen)
	copy(edek, data[offset+saltLen:offset+saltLen+wrappedDEKLen])

	payload := make([]byte, len(data)-headerLength)
	copy(payload, data[headerLength:])

	return EncryptedDataPayload{
		Salt:       Salt(salt),
		WrappedDEK: WrappedDEK(edek),
		Payload:    Ciphertext(payload),
	}, nil
}
