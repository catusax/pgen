package generator

import (
	"fmt"
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/catusax/pgen/generator/custom_func"
	"gopkg.in/yaml.v2"
)

// Config represents the configuration
type Config struct {
	// OnceFiles only generate once when creating new project
	OnceFiles []string `yaml:"once_files"`
	// SoftFiles will be generate every time when using generate command
	SoftFiles []string `yaml:"soft_files"`
	// HardFiles will be generate every time when using generate command with --hard flag
	HardFiles []string `yaml:"hard_files"`
	// DefaultENVs is default environment variables when creating new project.
	// on existing project, Makefile variables will replace these variables.
	DefaultENVs map[string]any `yaml:"default_envs"`

	WasmFuncs []custom_func.Libs `yaml:"wasm_funcs"`

	confDir string
}

var (
	c          *Config
	configOnce sync.Once
)

func Conf() *Config {
	configOnce.Do(initConfig)

	return c
}

func initConfig() {
	var err error
	c, err = readConfig()
	if err != nil {
		panic(err)
	}
}

func readConfig() (*Config, error) {
	config := Config{}
	fileBytes, readPath, err := ReadFile(filepath.Join(".template", ".pgen_config.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	log.Infoln("loaded config from file: ",
		filepath.Join(readPath, ".template", ".pgen_config.yaml"))

	err = yaml.Unmarshal(fileBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	config.confDir = filepath.Join(readPath, ".template")

	return &config, nil
}
