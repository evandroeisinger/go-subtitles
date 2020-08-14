package subtitles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	validPaths := []string{
		"testdata/sample.srt",
	}

	for _, path := range validPaths {
		subtitle, err := LoadFile(path)

		assert.Equal(t, Subtitle{}, subtitle, "should returns a subtitle")
		assert.Nil(t, err, "not returns error")
	}
}

func TestInvalidLoadFile(t *testing.T) {
	invalidPaths := []string{
		"testdata/invalid-path.srt",
		"testdata/invalid-file.srt",
	}

	for _, path := range invalidPaths {
		subtitle, err := LoadFile(path)

		assert.NotNil(t, err, "should returns a Error")
		assert.EqualError(t, err, fmt.Sprintf("Invalid file: %s", path))
		assert.Equal(t, Subtitle{}, subtitle, "should returns a empty subtitle")
	}
}
