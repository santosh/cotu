package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cotu",
		Usage: "developers swiss army knife",
		Commands: []*cli.Command{
			{
				Name:  "s3",
				Usage: "tasks related to aws s3",
				Subcommands: []*cli.Command{
					{
						Name:  "auth",
						Usage: "authenticate with aws",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("auth with aws in progress")
							return nil
						},
					},
					{
						Name:    "upload",
						Aliases: []string{"u"},
						Usage:   "upload to a bucket",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "bucket",
								Value: "cotu",
								Usage: "destination bucket",
							},
						},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("uploading to s3 in progress")
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
