package generator

import (
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration
type Config struct {
	// OnceFiles only generate once when creating new project
	OnceFiles []string `yaml:"onceFiles"`
	// SoftFiles will be generate every time when using generate command
	SoftFiles []string `yaml:"softFiles"`
	// HardFiles will be generate every time when using generate command with --hard flag
	HardFiles []string `yaml:"hardFiles"`
	// DefaultENVs is default environment variables when creating new project.
	// on existing project, Makefile variables will replace these variables.
	DefaultENVs map[string]any `yaml:"defaultEnvs"`
}

func ReadConfig() (*Config, error) {
	config := Config{}
	readPath, fileBytes, err := ReadFile(filepath.Join(".template", ".pgen_config.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	fmt.Println("loaded config from file: ", readPath) // TODO: log level

	err = yaml.Unmarshal(fileBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
