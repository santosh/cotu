package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/urfave/cli/v2"
)

var ListCommand cli.Command = cli.Command{
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
		ListBucket(path)
		return nil
	},
}

// ListBucket lists all the items in a bucket. A `path` can be pass to filter
// the results to that subdirectory only.
// Known issue: Subdirectories are not shown in the results. The way S3
// works, there is not concept of directories. Hierarchy is flat; all files
// are at same level.
func ListBucket(path string) {
	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("cotu"),
		Prefix: aws.String(path),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range output.Contents {
		fmt.Printf("%s\n", aws.ToString(object.Key))
	}
}
