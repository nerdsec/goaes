package internal

import "errors"

// Encrypt generates a new salt, creates the kek from the passphrase and the new
// salt, creates a new dek, wraps the dek with the kek, encrypts the data with
// the dek, then returns the edek, salt, and ciphertext.
//
// When built with GOEXPERIMENT=runtimesecret, all sensitive intermediates (KEK,
// DEK, decoded passphrase, and internal crypto buffers) are automatically zeroed
// by the runtime after secretDo returns. The manual Clear calls are retained as
// defense-in-depth for platforms without secret mode support.
func Encrypt(passphrase string, data []byte) (EncryptedDataPayload, error) {
	if len(passphrase) == 0 {
		return EncryptedDataPayload{}, errors.New("passphrase cannot be empty")
	}

	var result EncryptedDataPayload
	var retErr error

	SecretDo(func() {
		salt, err := NewSalt()
		if err != nil {
			retErr = err

			return
		}

		kek, err := NewKEKFromEnvB64(passphrase, salt)
		if err != nil {
			retErr = err

			return
		}
		defer Clear(kek)

		dek, err := NewDEK()
		if err != nil {
			retErr = err

			return
		}
		defer Clear(dek)

		edek, err := WrapDEK(dek, kek)
		if err != nil {
			retErr = err

			return
		}

		ct, err := EncryptData(data, dek)
		if err != nil {
			retErr = err

			return
		}

		result = EncryptedDataPayload{
			WrappedDEK: edek,
			Salt:       salt,
			Payload:    ct,
		}
	})

	return result, retErr
}
