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
	assert.EqualError(t, err, "Unsupported extension: .mp4")
}

func TestLoad(t *testing.T) {
	subtitle, err := Load("testdata/empty.srt")

	assert.Equal(t, "", subtitle.content)
	assert.Equal(t, 0, len(subtitle.blocks), "should have 0 blocks parsed")
	assert.Nil(t, err, "should not returns errors")
}
