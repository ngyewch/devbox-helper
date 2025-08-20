package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ngyewch/devbox-helper/devbox"
	"github.com/urfave/cli/v3"
)

func doProjectLatest(ctx context.Context, cmd *cli.Command) error {
	projectDir := cmd.String(projectDirFlag.Name)

	configPath := filepath.Join(projectDir, "devbox.json")

	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	config, err := devbox.ParseConfig(f)
	if err != nil {
		return err
	}

	client := devbox.NewClient(&http.Client{})

	for _, pkg := range config.Packages {
		resolveResponse, err := client.Resolve(devbox.ResolveRequest{
			Name:    pkg.Name,
			Version: "latest",
		})
		if err != nil {
			return err
		}
		if pkg.Version != resolveResponse.Version {
			fmt.Printf("%s@%s -> %s (latest)\n", pkg.Name, pkg.Version, resolveResponse.Version)
		} else {
			fmt.Printf("%s@%s (up-to-date)\n", pkg.Name, pkg.Version)
		}
	}

	return nil
}
