package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	subtitle, err := Load("testdata/empty.srt")

	assert.Equal(t, "", subtitle.content)
	assert.Equal(t, 0, len(subtitle.blocks), "should have 0 blocks parsed")
	assert.Nil(t, err, "should not returns errors")
}
