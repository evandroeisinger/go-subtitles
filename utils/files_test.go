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
	}

	for _, file := range fileTests {
		content, err := LoadFileContent(file.path)
		assert.Equal(t, file.content, content)
		assert.Nil(t, err)
	}
}

func TestLoadInvalidFileContent(t *testing.T) {
	filePaths := []string{
		"testdata/invalid.txt",
	}

	for _, path := range filePaths {
		content, err := LoadFileContent(path)
		assert.Equal(t, "", content)
		assert.NotNil(t, err)
	}
}
