package main

import (
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
					&s3.AuthCommand,
					&s3.UploadCommand,
					&s3.ListCommand,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
