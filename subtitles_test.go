package subtitles

import (
	"subtitles/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	filePath := "testdata/sample.srt"
	expectedContent := utils.LoadFileContent(filePath)

	subtitle, err := Load(filePath)

	assert.Equal(t, expectedContent, subtitle.content, "should have content loaded from file")
	assert.Equal(t, 5, len(subtitle.blocks), "should have 5 blocks parsed")
	assert.Nil(t, err, "should not returns errors")
}

func TestInvalidLoad(t *testing.T) {
	invalidFiles := []struct {
		path    string
		message string
	}{
		{path: "testdata/invalid-path.srt", message: "Invalid file: testdata/invalid-path.srt"},
		{path: "testdata/invalid-file.srt", message: "Invalid file: testdata/invalid-file.srt"},
		{path: "testdata/empty.srt", message: "Invalid file content: testdata/empty.srt\n"},
	}

	for _, file := range invalidFiles {
		subtitle, err := Load(file.path)

		assert.NotNil(t, err, "should returns a Error")
		assert.EqualError(t, err, file.message)
		assert.Equal(t, Subtitle{}, subtitle, "should returns a empty subtitle")
	}
}
