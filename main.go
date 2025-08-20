package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var (
	version string

	projectDirFlag = &cli.StringFlag{
		Name:    "project-dir",
		Usage:   "project directory",
		Value:   ".",
		Sources: cli.EnvVars("PROJECT_DIR"),
	}
)

func main() {
	app := &cli.Command{
		Name:    "devbox-helper",
		Usage:   "Devbox helper",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:  "project",
				Usage: "project",
				Flags: []cli.Flag{
					projectDirFlag,
				},
				Commands: []*cli.Command{
					{
						Name:   "latest",
						Usage:  "latest",
						Action: doProjectLatest,
					},
				},
			},
			{
				Name:      "latest",
				Usage:     "latest",
				ArgsUsage: "(package spec)",
				Action:    doLatest,
			},
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
