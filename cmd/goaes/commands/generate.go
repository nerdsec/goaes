package commands

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/nerdsec/goaes/internal"
	"github.com/urfave/cli/v3"
)

func Generate(ctx context.Context, cmd *cli.Command) error {
	var encoded string
	var retErr error

	internal.SecretDo(func() {
		key, err := internal.NewDEK()
		if err != nil {
			retErr = err

			return
		}
		defer internal.Clear(key)

		encoded = base64.StdEncoding.EncodeToString(key)
	})

	if retErr != nil {
		return retErr
	}

	fmt.Println(encoded)

	return nil
}
