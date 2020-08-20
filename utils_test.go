package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripBom(t *testing.T) {
	assert.Equal(t, "Lorem ipsum", StripBOM("\ufeffLorem ipsum"))
}

func TestIsEmptyFile(t *testing.T) {
	assert.True(t, IsEmptyFile("testdata/empty.srt"))
	assert.False(t, IsEmptyFile("testdata/sample.srt"))
}
