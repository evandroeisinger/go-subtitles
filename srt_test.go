package subtitles

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSubtitleWithInvalidBlockIndex(t *testing.T) {
	lines := []struct {
		content string
		err     string
	}{
		{BOMUnicode, "Invalid subtitle SRT: Expected index 1 at line 1 got: "},
		{"0", "Invalid subtitle SRT: Expected index 1 at line 1 got: 0"},
		{"2", "Invalid subtitle SRT: Expected index 1 at line 1 got: 2"},
		{"one", "Invalid subtitle SRT: Expected index 1 at line 1 got: one"},
	}

	for _, line := range lines {
		reader := strings.NewReader(line.content)
		subtitle, err := NewSRTParser().Parse(reader)

		assert.Nil(t, subtitle)
		assert.EqualError(t, err, line.err)
	}
}

func TestParseSubtitleWithInvalidBlockTime(t *testing.T) {
	lines := []struct {
		content string
		err     string
	}{
		{"1\nLorem ipsum", "Invalid subtitle SRT: Expected duration with pattern (hh:mm:ss,fff --> hh:mm:ss,fff) at line 2 got: Lorem ipsum"},
		{"1\n00:00:00", "Invalid subtitle SRT: Expected duration with pattern (hh:mm:ss,fff --> hh:mm:ss,fff) at line 2 got: 00:00:00"},
		{"1\n00:00:00 -> 99:99:99", "Invalid subtitle SRT: Expected duration with pattern (hh:mm:ss,fff --> hh:mm:ss,fff) at line 2 got: 00:00:00 -> 99:99:99"},
	}

	for _, line := range lines {
		reader := strings.NewReader(line.content)
		subtitle, err := NewSRTParser().Parse(reader)

		assert.Nil(t, subtitle)
		assert.EqualError(t, err, line.err)
	}
}

func TestParseSubtitleWithoutTextLines(t *testing.T) {
	reader := strings.NewReader("1\n00:00:00,000 --> 00:01:00,000")
	subtitle, err := NewSRTParser().Parse(reader)

	assert.Nil(t, subtitle)
	assert.EqualError(t, err, "Invalid subtitle SRT: Expected text at line 3 got: empty line")
}

func TestParseSimpleBlock(t *testing.T) {
	content := strings.NewReader("1\n00:00:00,000 --> 00:01:00,000\nLorem ipsum dolor sit amet\ndolor sit amet")
	subtitle, err := NewSRTParser().Parse(content)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(subtitle.Blocks))
	assert.Equal(t, "0s", subtitle.Blocks[0].StartAt.String())
	assert.Equal(t, "1m0s", subtitle.Blocks[0].FinishAt.String())
	assert.Equal(t, "Lorem ipsum dolor sit amet", subtitle.Blocks[0].Lines[0])
	assert.Equal(t, "dolor sit amet", subtitle.Blocks[0].Lines[1])
}

func TestParseSimpleBlockWithBOM(t *testing.T) {
	content := BOMUnicode + "1\n00:00:00,000 --> 00:01:00,000\nLorem ipsum dolor sit amet\ndolor sit amet"
	reader := strings.NewReader(content)
	subtitle, err := NewSRTParser().Parse(reader)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(subtitle.Blocks))
	assert.Equal(t, "0s", subtitle.Blocks[0].StartAt.String())
	assert.Equal(t, "1m0s", subtitle.Blocks[0].FinishAt.String())
	assert.Equal(t, "Lorem ipsum dolor sit amet", subtitle.Blocks[0].Lines[0])
	assert.Equal(t, "dolor sit amet", subtitle.Blocks[0].Lines[1])
}

func TestParse(t *testing.T) {
	file, _ := os.Open("testdata/sample.srt")
	defer file.Close()

	subtitle, err := NewSRTParser().Parse(file)

	blocks := []struct {
		startAt  string
		finishAt string
		lines    []string
	}{
		{
			startAt:  "0s",
			finishAt: "1s",
			lines: []string{
				"Lorem ipsum dolor sit amet",
			},
		},
		{
			startAt:  "1s",
			finishAt: "2s",
			lines: []string{
				"Consectetur adipiscing elit,",
				"sed do eiusmod tempor incididunt",
			},
		},
		{
			startAt:  "2s",
			finishAt: "3s",
			lines: []string{
				"<i>Ut labore et dolore magna aliqua<i>",
			},
		},
		{
			startAt:  "3.5s",
			finishAt: "4.5s",
			lines: []string{
				"Ut enim ad minim veniam,",
				"quis <b>nostrud exercitation</b> ullamco",
			},
		},
		{
			startAt:  "4.5s",
			finishAt: "5s",
			lines: []string{
				"Sed do eiusmod tempor incididunt,",
				"lorem ipsum dolor sit amet!",
				"Ut enim ad minim veniam",
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, len(blocks), len(subtitle.Blocks))

	for index, block := range blocks {
		assert.Equal(t, block.startAt, subtitle.Blocks[index].StartAt.String())
		assert.Equal(t, block.finishAt, subtitle.Blocks[index].FinishAt.String())
		assert.Equal(t, block.lines, subtitle.Blocks[index].Lines)
	}
}

func TestFormatEmptySubtitle(t *testing.T) {
	subtitle := NewSubtitle()
	content, err := NewSRTFormatter().Format(subtitle)

	assert.EqualError(t, err, "Invalid subtitle SRT: Empty blocks")
	assert.Equal(t, "", content)
}

func TestFormatSimpleBlock(t *testing.T) {
	block := NewBlock()
	block.StartAt, _ = time.ParseDuration("0s500ms")
	block.FinishAt, _ = time.ParseDuration("1m30s")
	block.Lines = []string{"Lorem ipsum dolor sit amet"}

	subtitle := NewSubtitle()
	subtitle.Blocks = append(subtitle.Blocks, block)

	content, err := NewSRTFormatter().Format(subtitle)
	assert.Nil(t, err)
	assert.Equal(t, "1\n00:00:00,500 --> 00:01:30,000\nLorem ipsum dolor sit amet\n", content)
}

func TestFormat(t *testing.T) {
	expectedContent, _ := ioutil.ReadFile("testdata/sample.srt")
	subtitle, _ := Load("testdata/sample.srt")

	content, err := NewSRTFormatter().Format(subtitle)
	assert.Equal(t, string(expectedContent), content)
	assert.Nil(t, err)
}
