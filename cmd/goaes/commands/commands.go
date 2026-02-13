package commands

import (
	"fmt"
	"os"
)

const (
	fileMode         = 0600
	passphraseEnvVar = "GOAES_PASSPHRASE"
	invalidArgsExit  = 2
	maxFileSize      = 1 << 30 // 1 GiB
)

var errSamePath = fmt.Errorf("source and destination must be different files")

func checkFileSize(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.Size() > maxFileSize {
		return fmt.Errorf("file too large: %d bytes (max %d)", info.Size(), maxFileSize)
	}

	return nil
}
