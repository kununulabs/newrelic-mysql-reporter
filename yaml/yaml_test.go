package yaml

import "testing"

func TestGetMetrics(t *testing.T) {
	workingFile := "./example.yaml"
	_, err := New(workingFile)
	if err != nil {
		t.Errorf("Could not read file %s!", workingFile)
	}

	nonExistingFile := "./nope-not-here.yaml"
	_, err = New(nonExistingFile)
	if err == nil {
		t.Errorf("Was able to read nonexisting file %s?", nonExistingFile)
	}
}
