package yaml

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// GetFile well, it gets the contents of a file
func GetFile(file string) ([]byte, error) {
	if len(file) < 1 {
		return nil, errors.New("no config file path given")
	}

	yaml, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return yaml, nil
}

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

// GetMetricsFromFile reads a metrics config from a yaml file
func GetMetricsFromFile(file string) ([]Metric, error) {
	yaml, err := GetFile(file)
	if err != nil {
		return []Metric{}, err
	}

	metrics, err := GetMetrics(yaml)
	if err != nil {
		return metrics, err
	}

	return metrics, nil
}

// GetMetrics returns all metrics defined in the config file
func GetMetrics(file []byte) ([]Metric, error) {
	config := Metrics{}

	if err := yaml.Unmarshal(file, &config); err != nil {
		return config.Metrics, err
	}

	if len(config.Metrics) < 1 {
		return config.Metrics, errors.New("no metrics found")
	}

	return config.Metrics, nil
}

// Attributes represents a full attributes config file
type Attributes struct {
	Attributes map[string]interface{} `yaml:"attributes"`
}

// GetAttributesFromFile returns parsed attributes from a yaml fiel
func GetAttributesFromFile(file string) (map[string]interface{}, error) {
	yaml, err := GetFile(file)
	if err != nil {
		// attributes file is optional
		if err.Error() == "no config file path given" {
			err = nil
		}
		return map[string]interface{}{}, err
	}

	attributes, err := GetAttributes(yaml)
	if err != nil {
		return attributes, err
	}

	return attributes, nil
}

// GetAttributes returns all attributes defined in the config file
func GetAttributes(file []byte) (map[string]interface{}, error) {
	config := Attributes{}

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return config.Attributes, err
	}

	if len(config.Attributes) < 1 {
		return config.Attributes, errors.New("no attributes found")
	}

	return config.Attributes, nil
}
