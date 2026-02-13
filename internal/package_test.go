package internal_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/nerdsec/goaes/internal"
)

func TestPackagePayload(t *testing.T) {
	//nolint:gosec // this is only for testing and not used for any implementation
	passphrase := []byte("dJyHOdMbG94EMvQGQrs6YZiXGiAGQgDYtx6+eqLufQg=")
	message := "hello"

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

func TestVersionRegression(t *testing.T) {
	passphrase := []byte(`kL9ntTh/DlgAAhu/AVYp17Qx9FaMjL7HZbpHK9hQyiE=`)

	tests := []struct {
		version   string
		payload   string
		plaintext string
	}{
		{
			version:   "v1.0.0",
			payload:   `R09BRVMBZxsY8Rtp6lCN+guCcX+o5T06kTJKd8TCgLjzRkHWYypNpXL2tNiXTRpWsy2JM4F8xi/ylhbrLmIwnTKxIz6FtwUQK/Gd0adOY9cRh/Fxz773sNpm9W9fe1vaFEJJJZWbkD76tphS1difZGhwtNe7FUOnexOZeewStcuvXwGt`,
			plaintext: "v1.0.0",
		},
	}

	for _, test := range tests {
		t.Run(test.version, (func(tt *testing.T) {
			raw, err := base64.StdEncoding.DecodeString(test.payload)
			if err != nil {
				tt.Error(err)
			}

			unpackaged, err := internal.UnpackagePayload(raw)
			if err != nil {
				tt.Fatal("unexpected error unpackaging payload")
			}

			plaintext, err := internal.Decrypt(
				passphrase,
				unpackaged.WrappedDEK,
				unpackaged.Payload,
				unpackaged.Salt,
			)
			if err != nil {
				tt.Fatal("failed to decrypt", err)
			}

			if !bytes.Equal(plaintext, []byte(test.plaintext)) {
				tt.Error("plaintext didn't match")
			}
		}))
	}
}
