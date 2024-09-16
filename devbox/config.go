package devbox

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Config struct {
	Name        string
	Description string
	Packages    []PackageSpec
	Env         map[string]string
	// TODO
}

type PackageSpec struct {
	Name    string
	Version string
	// TODO
}

type configFromSchema struct {
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Packages    any               `json:"packages,omitempty"`
	Env         map[string]string `json:"env,omitempty"`
	Shell       any               `json:"shell,omitempty"`
}

func ParseConfig(r io.Reader) (*Config, error) {
	jsonDecoder := json.NewDecoder(r)
	var internalConfig configFromSchema
	err := jsonDecoder.Decode(&internalConfig)
	if err != nil {
		return nil, err
	}

	config := Config{
		Name:        internalConfig.Name,
		Description: internalConfig.Description,
		Env:         internalConfig.Env,
	}

	switch v := internalConfig.Packages.(type) {
	case []string:
		for _, packageSpec := range v {
			parts := strings.SplitN(packageSpec, "@", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid package spec: %s", packageSpec)
			}
			config.Packages = append(config.Packages, PackageSpec{
				Name:    parts[0],
				Version: parts[1],
			})
		}
	case map[string]any:
		for packageName, packageInfo := range v {
			pkg := PackageSpec{
				Name: packageName,
			}
			switch v2 := packageInfo.(type) {
			case string:
				pkg.Version = v2
			case map[string]any:
				for key, value := range v2 {
					switch key {
					case "version":
						pkg.Version = value.(string)
					default:
						// TODO
					}
				}
			}
			config.Packages = append(config.Packages, pkg)
		}
	default:
		return nil, fmt.Errorf("unsupported package type: %T", v)
	}

	return &config, nil
}
