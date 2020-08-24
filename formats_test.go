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
	assert.EqualError(t, err, "Invalid file invalid.mp4: Parser for extension not found")
}

func TestFormatterForFileSRT(t *testing.T) {
	formatter, err := FormatterForFile("valid.srt")

	assert.Equal(t, NewSRTFormatter(), formatter)
	assert.Nil(t, err)
}

func TestFormatterForInvalidFile(t *testing.T) {
	formatter, err := FormatterForFile("invalid.mp4")

	assert.Nil(t, formatter)
	assert.EqualError(t, err, "Invalid file invalid.mp4: Formatter for extension not found")
}
