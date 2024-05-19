package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config represents the structure of the configuration file.
type Config struct {
	Services         []Service `yaml:"services"`
	Timeout          int       `yaml:"timeout"`
	ConcurrencyLevel int       `yaml:"concurrencyLevel"`
	OutputFormat     string    `yaml:"outputFormat"`
	OutputPath       string    `yaml:"outputPath"`
}

// Service represents a single service to check.
type Service struct {
	Address string `yaml:"address"`
}

// LoadConfig loads the configuration from a YAML file.
func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
