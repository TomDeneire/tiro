package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// General functions

func ReadFile(path string) (string, error) {
	result := ""
	file, _ := os.Open(path)
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return result, err
	}
	result = string(b)
	return result, nil
}

// Test functions

// Compare result and expected for tests
func Check(result string, expected string, t *testing.T) {
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}
