package subtitles

import (
	"subtitles/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromFile(t *testing.T) {
	filePath := "testdata/sample.srt"
	expectedContent := utils.LoadFileContent(filePath)

	expectedSubtitle := Subtitle{
		content: expectedContent,
		blocks:  []Block{},
	}

	subtitle, err := LoadFromFile(filePath)

	assert.Equal(t, expectedSubtitle, subtitle, "should returns a subtitle")
	assert.Nil(t, err, "not returns error")
}

func TestInvalidLoadFromFile(t *testing.T) {
	invalidFiles := []struct {
		path    string
		message string
	}{
		{path: "testdata/invalid-path.srt", message: "Invalid file: testdata/invalid-path.srt"},
		{path: "testdata/invalid-file.srt", message: "Invalid file: testdata/invalid-file.srt"},
		{path: "testdata/empty.srt", message: "Invalid file content: testdata/empty.srt\n"},
	}

	for _, file := range invalidFiles {
		subtitle, err := LoadFromFile(file.path)

		assert.NotNil(t, err, "should returns a Error")
		assert.EqualError(t, err, file.message)
		assert.Equal(t, Subtitle{}, subtitle, "should returns a empty subtitle")
	}
}
