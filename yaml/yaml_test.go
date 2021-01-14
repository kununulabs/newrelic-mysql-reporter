package yaml

import "testing"

func TestGetMetrics(t *testing.T) {
	workingFile := "./metrics-example.yaml"
	_, err := GetMetricsFromFile(workingFile)
	if err != nil {
		t.Errorf("Could not read file %s!", workingFile)
	}

	nonExistingFile := "./nope-not-here.yaml"
	_, err = GetMetricsFromFile(nonExistingFile)
	if err == nil {
		t.Errorf("Was able to read nonexisting file %s?", nonExistingFile)
	}

	garbledYaml := "ajkasljklfjalkfjaklfjklajk"
	metrics, err := GetMetrics([]byte(garbledYaml))
	if err == nil {
		t.Log(metrics)
		t.Errorf("Was able to parse random data into metrics?")
	}
}

func TestGetAttributes(t *testing.T) {
	workingFile := "./attributes-example.yaml"
	_, err := GetAttributesFromFile(workingFile)
	if err != nil {
		t.Errorf("Could not read file %s!", workingFile)
	}

	nonExistingFile := "./nope-not-here.yaml"
	_, err = GetAttributesFromFile(nonExistingFile)
	if err == nil {
		t.Errorf("Was able to read nonexisting file %s?", nonExistingFile)
	}

	garbledYaml := "ajkasljklfjalkfjaklfjklajk"
	_, err = GetAttributes([]byte(garbledYaml))
	if err == nil {
		t.Errorf("Was able to parse random data into attributes?")
	}
}
