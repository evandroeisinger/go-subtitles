package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	fileTests := []struct {
		path   string
		exists bool
	}{
		{path: "../testdata/sample.srt", exists: true},
		{path: "../testdata/invalid.srt", exists: false},
	}

	for _, file := range fileTests {
		assert.Equal(t, file.exists, FileExists(file.path), file.path)
	}
}
