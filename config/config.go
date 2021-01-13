package config

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Metric represents a single metric to process
type Metric struct {
	Comment string `yaml:"comment"`
	Name    string `yaml:"name"`
	Query   string `yaml:"query"`
}

// Metrics represents the full metrics config file
type Metrics struct {
	Metrics []Metric `yaml:"metrics"`
}

// Attributes represents a full attributes config file
type Attributes struct {
	Attributes map[string]interface{} `yaml:"attributes"`
}

// GetMetrics returns all metrics defined in the config file
func GetMetrics(configFile string) ([]Metric, error) {
	config := Metrics{}

	if len(configFile) < 1 {
		return config.Metrics, errors.New("no config file path given")
	}

	configYaml, err := ioutil.ReadFile(configFile)
	if err != nil {
		return config.Metrics, err
	}

	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		return config.Metrics, err
	}

	return config.Metrics, nil
}

// GetAttributes returns all attributes defined in the config file
func GetAttributes(configFile string) (map[string]interface{}, error) {
	config := Attributes{}

	if len(configFile) < 1 {
		return config.Attributes, errors.New("no config file path given")
	}

	configYaml, err := ioutil.ReadFile(configFile)
	if err != nil {
		return config.Attributes, err
	}

	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		return config.Attributes, err
	}

	return config.Attributes, nil
}
