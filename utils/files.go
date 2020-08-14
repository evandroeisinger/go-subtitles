package utils

import (
	"io/ioutil"
	"os"
)

// FileExists method
func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

// LoadFileContent method
func LoadFileContent(path string) (string, error) {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
