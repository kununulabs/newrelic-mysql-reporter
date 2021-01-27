package yaml

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config represents the full config file
type Config struct {
	Metrics    []Metric               `yaml:"metrics"`
	Attributes map[string]interface{} `yaml:"attributes"`
}

// Metric represents a single metric to process
type Metric struct {
	Comment string `yaml:"comment"`
	Name    string `yaml:"name"`
	Query   string `yaml:"query"`
}

// New unmarshals a file to our config format
func New(file string) (*Config, error) {
	if len(file) < 1 {
		return nil, errors.New("no config file path given")
	}

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := yaml.Unmarshal(raw, config); err != nil {
		return nil, err
	}

	return config, nil
}
