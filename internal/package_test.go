package internal_test

import (
	"bytes"
	"testing"

	"github.com/nerdsec/goaes/internal"
)

func TestPackagePayload(t *testing.T) {
	const (
		//nolint:gosec // this is only for testing and not used for any implementation
		passphrase = "dJyHOdMbG94EMvQGQrs6YZiXGiAGQgDYtx6+eqLufQg="
		message    = "hello"
	)

	payload, err := internal.Encrypt(passphrase, []byte(message))
	if err != nil {
		t.Fatal("failed to encrypt payload during test")
	}

	packaged, err := internal.PackagePayload(payload)
	if err != nil {
		t.Fatal("failed to package payload during test")
	}

	if packaged == nil {
		t.Fatal("package shouldn't be nil")
	}

	unpackaged, err := internal.UnpackagePayload(packaged)
	if err != nil {
		t.Fatal("unexpected error unpackaging payload")
	}

	plaintext, err := internal.Decrypt(
		passphrase,
		unpackaged.WrappedDEK,
		unpackaged.Payload,
		unpackaged.Salt,
	)
	if err != nil {
		t.Fatal("failed to decrypt")
	}

	if !bytes.Equal(plaintext, []byte(message)) {
		t.Error("plaintext didn't match")
	}
}

func TestUnpackagePayloadTooShort(t *testing.T) {
	_, err := internal.UnpackagePayload([]byte("short"))
	if err == nil {
		t.Error("should fail with data too short")
	}
}

func TestUnpackagePayloadBadMagic(t *testing.T) {
	// Create a buffer that's long enough but has wrong magic bytes
	data := make([]byte, 100)
	copy(data, []byte("WRONG"))

	_, err := internal.UnpackagePayload(data)
	if err == nil {
		t.Error("should fail with invalid magic bytes")
	}
}

func TestUnpackagePayloadBadVersion(t *testing.T) {
	// Create a buffer with correct magic but wrong version
	data := make([]byte, 100)
	copy(data, []byte("GOAES"))
	data[5] = 0xFF

	_, err := internal.UnpackagePayload(data)
	if err == nil {
		t.Error("should fail with unsupported version")
	}
}
