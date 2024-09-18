package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	projectDirFlag = &cli.PathFlag{
		Name:    "project-dir",
		Usage:   "project directory",
		Value:   ".",
		EnvVars: []string{"PROJECT_DIR"},
	}
)

func main() {
	app := &cli.App{
		Name:  "devbox-helper",
		Usage: "Devbox helper",
		Commands: []*cli.Command{
			{
				Name:  "project",
				Usage: "project",
				Flags: []cli.Flag{
					projectDirFlag,
				},
				Subcommands: []*cli.Command{
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
