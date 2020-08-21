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

func TestOpenInvalidFile(t *testing.T) {
	files := []struct {
		path string
		err  string
	}{
		{"testdata/invalid.srt", "Invalid file testdata/invalid.srt: File not exist"},
		{"testdata/empty.srt", "Invalid file testdata/empty.srt: Empty file"},
		{"testdata", "Invalid file testdata: Its not a file"},
	}

	for _, file := range files {
		f, err := OpenFile(file.path)

		assert.Nil(t, f)
		assert.EqualError(t, err, file.err)
	}
}

func TestCreateInvalidFile(t *testing.T) {
	f, err := CreateFile("testdata/sample.srt")

	assert.Nil(t, f)
	assert.EqualError(t, err, "Invalid file testdata/sample.srt: File already exist")
}
