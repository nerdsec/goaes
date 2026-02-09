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
		t.Error("failed to encrypt payload during test")
	}

	packaged := internal.PackagePayload(payload)
	if packaged == nil {
		t.Error("package shouldn't be nil")
	}

	unpackaged, err := internal.UnpackagePayload(packaged)
	if err != nil {
		t.Error("unexpected error unpackaging payload")
	}

	plaintext, err := internal.Decrypt(
		passphrase,
		unpackaged.DEK,
		unpackaged.Payload,
		unpackaged.Salt,
	)
	if err != nil {
		t.Error("failed to decrypt")
	}

	if !bytes.Equal(plaintext, []byte(message)) {
		t.Error("plaintext didn't match")
	}
}
