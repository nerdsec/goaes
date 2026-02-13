package commands

import (
	"context"
	"fmt"

	"github.com/nerdsec/goaes/internal"
	"github.com/urfave/cli/v3"
)

func Features(ctx context.Context, cmd *cli.Command) error {
	fmt.Printf("Go secret/runtime: %v\n", internal.SecretEnabled())

	return nil
}
