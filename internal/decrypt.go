package internal

// Decrypt recreates the kek from a passphrase and a salt, unwraps the dek using
// the kek, decrypts the data using the dek, and then returns the plaintext.
func Decrypt(passphrase string, edek WrappedDEK, ct Ciphertext, salt Salt) ([]byte, error) {
	kek, err := NewKEKFromEnvB64(passphrase, salt)
	if err != nil {
		return nil, err
	}

	dek, err := UnwrapDEK(edek, kek)
	if err != nil {
		return nil, err
	}

	pt, err := DecryptData(ct, dek)
	if err != nil {
		return nil, err
	}

	return pt, nil
}
