package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	subtitle, err := LoadFile("testdata/sample.srt")

	assert.Equal(t, subtitle, Subtitle{}, "should returns a subtitle")
	assert.Equal(t, err, nil, "not returns error")
}
