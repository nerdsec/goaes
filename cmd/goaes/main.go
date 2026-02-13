package main

import (
	"context"
	"log"
	"os"

	"github.com/nerdsec/goaes/cmd/goaes/commands"
	"github.com/nerdsec/goaes/internal"
	"github.com/urfave/cli/v3"
)

var (
	version = "dev"
)

func main() {
	cmd := &cli.Command{
		Name:    "goaes",
		Usage:   "Simple AES encryption built with Go",
		Version: version,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return cli.DefaultShowRootCommandHelp(cmd)
		},
		Metadata: map[string]interface{}{
			"secret.mode": internal.SecretEnabled(),
		},
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generate a base64 encoded key",
				Action:  commands.Generate,
			},
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "Encrypt a file",
				Action:  commands.Encrypt,
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "source",
					},
					&cli.StringArg{
						Name: "destination",
					},
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "Decrypt a file",
				Action:  commands.Decrypt,
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "source",
					},
					&cli.StringArg{
						Name: "destination",
					},
				},
			},
			{
				Name:    "features",
				Aliases: []string{"f"},
				Usage:   "Show features enabled",
				Action:  commands.Features,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
