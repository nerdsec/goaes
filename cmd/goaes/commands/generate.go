package commands

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/nerdsec/goaes/internal"
	"github.com/urfave/cli/v3"
)

func Generate(ctx context.Context, cmd *cli.Command) error {
	var retErr error

	internal.SecretDo(func() {
		key, err := internal.NewDEK()
		if err != nil {
			retErr = err

			return
		}
		defer internal.Clear(key)

		encoded := make([]byte, base64.StdEncoding.EncodedLen(len(key)))
		base64.StdEncoding.Encode(encoded, key)
		defer internal.Clear(encoded)

		fmt.Println(string(encoded))
	})

	return retErr
}
