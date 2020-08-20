package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserForFileSRT(t *testing.T) {
	parser, err := ParserForFile("valid.srt")

	assert.Equal(t, NewSRTParser(), parser)
	assert.Nil(t, err)
}

func TestParserForInvalidFile(t *testing.T) {
	parser, err := ParserForFile("invalid.mp4")

	assert.Nil(t, parser)
	assert.EqualError(t, err, "Invalid file invalid.mp4: Extension not supported")
}

func TestLoadInvalidFile(t *testing.T) {
	files := []struct {
		path string
		err  string
	}{
		{"testdata/unsupported.uspd", "Invalid file testdata/unsupported.uspd: Extension not supported"},
		{"testdata/invalid.srt", "Invalid file testdata/invalid.srt: File not exist"},
		{"testdata/empty.srt", "Invalid file testdata/empty.srt: Empty file"},
		{"testdata", "Invalid file testdata: Its not a file"},
	}

	for _, file := range files {
		subtitle, err := Load(file.path)

		assert.Nil(t, subtitle)
		assert.EqualError(t, err, file.err)
	}
}

func TestLoadFile(t *testing.T) {
	subtitle, err := Load("testdata/sample.srt")

	assert.Equal(t, 5, len(subtitle.Blocks))
	assert.Nil(t, err)
}
