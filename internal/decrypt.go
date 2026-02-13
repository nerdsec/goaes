package internal

import "errors"

// Decrypt recreates the kek from a passphrase and a salt, unwraps the dek using
// the kek, decrypts the data using the dek, and then returns the plaintext.
//
// When built with GOEXPERIMENT=runtimesecret, all sensitive intermediates (KEK,
// DEK, decoded passphrase, and internal crypto buffers) are automatically zeroed
// by the runtime after secretDo returns. The manual Clear calls are retained as
// defense-in-depth for platforms without secret mode support.
func Decrypt(passphrase string, edek WrappedDEK, ct Ciphertext, salt Salt) ([]byte, error) {
	if len(passphrase) == 0 {
		return nil, errors.New("passphrase cannot be empty")
	}

	if len(edek) == 0 {
		return nil, errors.New("wrapped DEK cannot be empty")
	}

	if len(ct) == 0 {
		return nil, errors.New("ciphertext cannot be empty")
	}

	if len(salt) == 0 {
		return nil, errors.New("salt cannot be empty")
	}

	var pt []byte
	var retErr error

	SecretDo(func() {
		kek, err := NewKEKFromEnvB64(passphrase, salt)
		if err != nil {
			retErr = err

			return
		}
		defer Clear(kek)

		dek, err := UnwrapDEK(edek, kek)
		if err != nil {
			retErr = err

			return
		}
		defer Clear(dek)

		plaintext, err := DecryptData(ct, dek)
		if err != nil {
			retErr = err

			return
		}

		pt = plaintext
	})

	return pt, retErr
}
