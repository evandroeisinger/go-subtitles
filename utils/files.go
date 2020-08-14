package utils

import (
	"os"
)

// FileExists method
func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}
