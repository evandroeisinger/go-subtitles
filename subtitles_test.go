package subtitles

import (
	"os"
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

func TestLoadInvalidFile(t *testing.T) {
	files := []struct {
		path string
		err  string
	}{
		{"testdata/unsupported.uspd", "Invalid file testdata/unsupported.uspd: Parser for extension not found"},
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

func TestWriteInvalidSubtitles(t *testing.T) {
	sampleSubtitle, _ := Load("testdata/sample.srt")

	files := []struct {
		path     string
		subtitle *Subtitle
		err      string
	}{
		{"empty.srt", NewSubtitle(), "Invalid subtitle SRT: Empty blocks"},
		{"subtitle.uspd", sampleSubtitle, "Invalid file subtitle.uspd: Formatter for extension not found"},
		{"testdata/sample.srt", sampleSubtitle, "Invalid file testdata/sample.srt: File already exist"},
	}

	for _, file := range files {
		content, err := Write(file.subtitle, file.path)

		assert.EqualError(t, err, file.err)
		assert.Equal(t, "", content)
	}
}

func TestWrite(t *testing.T) {
	sampleSubtitle, _ := Load("testdata/sample.srt")

	content, err := Write(sampleSubtitle, "testdata/tmp.srt")
	assert.Nil(t, err)
	assert.Equal(t, 444, len(content))

	tmpSubtitle, _ := Load("testdata/tmp.srt")
	assert.EqualValues(t, sampleSubtitle, tmpSubtitle)

	os.Remove("testdata/tmp.srt")
}
