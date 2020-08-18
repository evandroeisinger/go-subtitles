package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatForFileSRT(t *testing.T) {
	format, err := FormatForFile("valid.srt")

	assert.Equal(t, NewSRT(), format)
	assert.Nil(t, err)
}

func TestFormatForInvalidFileFormat(t *testing.T) {
	format, err := FormatForFile("invalid.mp4")

	assert.Nil(t, format)
	assert.EqualError(t, err, "Unsupported extension: .mp4")
}

func TestLoad(t *testing.T) {
	subtitle, err := Load("testdata/empty.srt")

	assert.Equal(t, "", subtitle.content)
	assert.Equal(t, 0, len(subtitle.blocks), "should have 0 blocks parsed")
	assert.Nil(t, err, "should not returns errors")
}
