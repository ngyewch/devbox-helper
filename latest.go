package main

import (
	"fmt"
	"github.com/ngyewch/devbox-helper/devbox"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"path/filepath"
)

func doLatest(cCtx *cli.Context) error {
	configPath := filepath.Join(".", "devbox.json")
	if cCtx.NArg() > 0 {
		path := cCtx.Args().Get(0)
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			configPath = filepath.Join(path, "devbox.json")
		} else {
			configPath = path
		}
	}

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
