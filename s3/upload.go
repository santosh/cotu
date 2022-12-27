package s3

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var UploadCommand cli.Command = cli.Command{
	Name:    "upload",
	Aliases: []string{"u"},
	Usage:   "upload to a bucket",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "path",
			Value: "cotu",
			Usage: "destination path in cotu bucket",
		},
	},
	Action: func(cCtx *cli.Context) error {
		fmt.Println("uploading to s3 in progress")
		return nil
	},
}

func UploadFileToBucket(ctx *cli.Context, file *os.File, prefix string) {
	client.PutObject(ctx)
}
