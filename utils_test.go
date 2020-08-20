package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripBom(t *testing.T) {
	assert.Equal(t, "Lorem ipsum", StripBOM("\ufeffLorem ipsum"))
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile("testdata/sample.srt"))
	assert.False(t, IsFile("testdata/"))
}

func TestIsEmptyFile(t *testing.T) {
	assert.True(t, IsEmptyFile("testdata/empty.srt"))
	assert.False(t, IsEmptyFile("testdata/sample.srt"))
}

func TestFileExist(t *testing.T) {
	assert.True(t, FileExist("testdata/sample.srt"))
	assert.False(t, FileExist("testdata/invalid.srt"))
}
