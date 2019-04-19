package util

import (
	"bytes"
	"testing"
)

func TestDumpTo(t *testing.T) {
	inputs := []string{"this", "is", "a", "list", "of", "test", "strings"}
	expectedResult := "this\nis\na\nlist\nof\ntest\nstrings\n"

	buf := bytes.NewBuffer(nil)
	DumpTo(inputs, buf)
	actualResult := buf.String()

	if actualResult != expectedResult {
		t.Errorf("Expected \"%s\", got \"%s\"", expectedResult, actualResult)
	}
}
