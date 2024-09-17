package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "devbox-helper",
		Usage: "Devbox helper",
		Commands: []*cli.Command{
			{
				Name:      "latest",
				Usage:     "latest",
				ArgsUsage: "[(directory with or path to devbox.json)]",
				Action:    doLatest,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
