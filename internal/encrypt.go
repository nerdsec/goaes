package internal

import "errors"

// Encrypt generates a new salt, creates the kek from the passphrase and the new
// salt, creates a new dek, wraps the dek with the kek, encrypts the data with
// the dek, then returns the edek, salt, and ciphertext.
func Encrypt(passphrase string, data []byte) (EncryptedDataPayload, error) {
	if len(passphrase) == 0 {
		return EncryptedDataPayload{}, errors.New("passphrase cannot be empty")
	}

	salt, err := NewSalt()
	if err != nil {
		return EncryptedDataPayload{}, err
	}

	kek, err := NewKEKFromEnvB64(passphrase, salt)
	if err != nil {
		return EncryptedDataPayload{}, err
	}
	defer clear(kek)

	dek, err := NewDEK()
	if err != nil {
		return EncryptedDataPayload{}, err
	}
	defer clear(dek)

	edek, err := WrapDEK(dek, kek)
	if err != nil {
		return EncryptedDataPayload{}, err
	}

	ct, err := EncryptData(data, dek)
	if err != nil {
		return EncryptedDataPayload{}, err
	}

	return EncryptedDataPayload{
		DEK:     edek,
		Salt:    salt,
		Payload: ct,
	}, nil
}
