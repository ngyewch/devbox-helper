package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	version         string
	commit          string
	commitTimestamp string

	projectDirFlag = &cli.PathFlag{
		Name:    "project-dir",
		Usage:   "project directory",
		Value:   ".",
		EnvVars: []string{"PROJECT_DIR"},
	}
)

func main() {
	app := &cli.App{
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

	cli.VersionPrinter = func(cCtx *cli.Context) {
		var parts []string
		if version != "" {
			parts = append(parts, fmt.Sprintf("version=%s", version))
		}
		if commit != "" {
			parts = append(parts, fmt.Sprintf("commit=%s", commit))
		}
		if commitTimestamp != "" {
			formattedCommitTimestamp := func(commitTimestamp string) string {
				epochSeconds, err := strconv.ParseInt(commitTimestamp, 10, 64)
				if err != nil {
					return ""
				}
				t := time.Unix(epochSeconds, 0)
				return t.Format(time.RFC3339)
			}(commitTimestamp)
			if formattedCommitTimestamp != "" {
				parts = append(parts, fmt.Sprintf("commitTimestamp=%s", formattedCommitTimestamp))
			}
		}
		fmt.Println(strings.Join(parts, " "))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
