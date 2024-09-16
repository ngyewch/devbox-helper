package main

import (
	"fmt"
	"github.com/ngyewch/devbox-outdated/devbox"
	"net/http"
	"os"
)

func main() {
	err := doMain(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

func doMain(args []string) error {
	f, err := os.Open(args[0])
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
		resolveRespose, err := client.Resolve(devbox.ResolveRequest{
			Name:    pkg.Name,
			Version: pkg.Version,
		})
		if err != nil {
			return err
		}
		if pkg.Version != resolveRespose.Version {
			fmt.Printf("%s@%s -> %s (latest)\n", pkg.Name, pkg.Version, resolveRespose.Version)
		} else {
			fmt.Printf("%s@%s (up-to-date)\n", pkg.Name, pkg.Version)
		}
	}

	return nil
}
