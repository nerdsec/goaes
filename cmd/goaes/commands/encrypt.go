package commands

import (
	"context"
	"os"
	"path/filepath"

	"github.com/nerdsec/goaes/internal"
	"github.com/urfave/cli/v3"
)

func Encrypt(ctx context.Context, cmd *cli.Command) error {
	source := cmd.StringArg("source")
	destination := cmd.StringArg("destination")

	if source == "" {
		return cli.Exit("missing source file", invalidArgsExit)
	}

	if destination == "" {
		destination = source + ".goaes"
	}

	passphrase := []byte(os.Getenv(passphraseEnvVar))
	if len(passphrase) == 0 {
		return cli.Exit("GOAES_PASSPHRASE environment variable is not set", 1)
	}

	defer internal.Clear(passphrase)

	source = filepath.Clean(source)
	destination = filepath.Clean(destination)

	if source == destination {
		return errSamePath
	}

	if err := checkFileSize(source); err != nil {
		return err
	}

	plaintext, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	payload, err := internal.Encrypt(passphrase, plaintext)
	if err != nil {
		return err
	}

	buffer, err := internal.PackagePayload(payload)
	if err != nil {
		return err
	}

	return os.WriteFile(destination, buffer, fileMode)
}
