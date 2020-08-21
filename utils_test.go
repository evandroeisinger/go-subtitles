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

func TestPathExist(t *testing.T) {
	assert.True(t, PathExist("testdata/sample.srt"))
	assert.False(t, PathExist("testdata/invalid.srt"))
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
	files := []struct {
		path string
		err  string
	}{
		{"testdata/sample.srt", "Invalid file testdata/sample.srt: File already exist"},
		{"tmp/invalid.srt", "Invalid file tmp/invalid.srt: File path not exist"},
	}

	for _, file := range files {
		f, err := CreateFile(file.path)

		assert.Nil(t, f)
		assert.EqualError(t, err, file.err)
	}
}
