package commands

import (
	"context"
	"os"
	"path/filepath"

	"github.com/nerdsec/goaes/internal"
	"github.com/urfave/cli/v3"
)

func Decrypt(ctx context.Context, cmd *cli.Command) error {
	source := cmd.StringArg("source")
	destination := cmd.StringArg("destination")

	if source == "" {
		return cli.Exit("missing source", invalidArgsExit)
	}

	if destination == "" {
		return cli.Exit("missing destination", invalidArgsExit)
	}

	passphrase := []byte(os.Getenv(passphraseEnvVar))
	if len(passphrase) == 0 {
		return cli.Exit("GOAES_PASSPHRASE environment variable is not set", 1)
	}

	defer internal.Clear(passphrase)

	source = filepath.Clean(source)
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	encryptedPayload, err := internal.UnpackagePayload(data)
	if err != nil {
		return err
	}

	plaintext, err := internal.Decrypt(
		passphrase,
		encryptedPayload.WrappedDEK,
		encryptedPayload.Payload,
		encryptedPayload.Salt,
	)
	if err != nil {
		return err
	}

	destination = filepath.Clean(destination)

	return os.WriteFile(destination, plaintext, fileMode)
}
