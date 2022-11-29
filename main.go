package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/santosh/cotu/s3"
	"github.com/urfave/cli/v2"
)

func init() {
	godotenv.Load()
}

func main() {
	app := &cli.App{
		Name:    "cotu",
		Usage:   "developers swiss army knife",
		Version: "0.1.0",
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
								Name:  "path",
								Value: "cotu",
								Usage: "destination path in cotu bucket",
							},
						},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("uploading to s3 in progress")
							return nil
						},
					},
					{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "list cotu bucket",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "path",
								Aliases:     []string{"p"},
								Value:       "",
								DefaultText: "\"\"",
								Usage:       "possible `SUBDIR` to list",
							},
						},
						Action: func(cCtx *cli.Context) error {
							path := cCtx.String("path")
							s3.ListBucket(path)
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
