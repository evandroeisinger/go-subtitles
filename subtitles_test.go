package subtitles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	fileTests := []struct {
		path string
	}{
		{path: "testdata/sample.srt"},
	}

	for _, file := range fileTests {
		subtitle, err := LoadFile(file.path)

		assert.Equal(t, subtitle, Subtitle{}, "should returns a subtitle")
		assert.Nil(t, err, "not returns error")
	}
}

func TestInvalidLoadFile(t *testing.T) {
	invalidFileTests := []struct {
		path string
	}{
		{path: "testdata/invalid-path.srt"},
		{path: "testdata/invalid-file.srt"},
	}

	for _, file := range invalidFileTests {
		subtitle, err := LoadFile(file.path)

		assert.NotNil(t, err, "should returns a Error")
		assert.EqualError(t, err, fmt.Sprintf("Invalid file: %s", file.path))
		assert.Equal(t, subtitle, Subtitle{}, "should returns a empty subtitle")
	}
}
