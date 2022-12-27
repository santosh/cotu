package s3

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var AuthCommand cli.Command = cli.Command{
	Name:  "auth",
	Usage: "authenticate with aws",
	Action: func(cCtx *cli.Context) error {
		fmt.Println("auth with aws in progress")
		return nil
	},
}
