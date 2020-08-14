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
		{path: "testdata/empty.txt", exists: true},
		{path: "testdata/invalid.srt", exists: false},
	}

	for _, file := range fileTests {
		assert.Equal(t, file.exists, FileExists(file.path), file.path)
	}
}

func TestLoadFileContent(t *testing.T) {
	fileTests := []struct {
		path    string
		content string
	}{
		{path: "testdata/empty.txt", content: ""},
		{path: "testdata/lorem.txt", content: "Lorem ipsum dolor sit amet"},
		{path: "testdata/invalid.txt", content: ""},
	}

	for _, file := range fileTests {
		assert.Equal(t, file.content, LoadFileContent(file.path))
	}
}
