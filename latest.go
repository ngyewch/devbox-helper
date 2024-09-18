package main

import (
	"fmt"
	"github.com/ngyewch/devbox-helper/devbox"
	"github.com/urfave/cli/v2"
	"net/http"
	"strings"
)

func doLatest(cCtx *cli.Context) error {
	packageSpec := cCtx.Args().Get(0)

	parts := strings.SplitN(packageSpec, "@", 2)
	packageName := parts[0]
	packageVersion := "latest"
	if len(parts) == 2 {
		packageVersion = parts[1]
	}

	client := devbox.NewClient(&http.Client{})

	resolveResponse, err := client.Resolve(devbox.ResolveRequest{
		Name:    packageName,
		Version: packageVersion,
	})
	if err != nil {
		return err
	}
	fmt.Printf("latest version: %s@%s\n", resolveResponse.Name, resolveResponse.Version)

	return nil
}
