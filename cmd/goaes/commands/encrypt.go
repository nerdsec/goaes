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

	passphrase := os.Getenv(passphraseEnvVar)
	if passphrase == "" {
		return cli.Exit("GOAES_PASSPHRASE environment variable is not set", 1)
	}

	source = filepath.Clean(source)
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

	destination = filepath.Clean(destination)

	return os.WriteFile(destination, buffer, fileMode)
}
