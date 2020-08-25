package subtitles

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	shiftedSubtitle, _ := Load("testdata/shifted_sample.srt")
	subtitle, _ := Load("testdata/sample.srt")

	shiftDuration, _ := time.ParseDuration("-500ms")
	subtitle.Shift(shiftDuration)

	assert.EqualValues(t, shiftedSubtitle, subtitle)
}

func TestCut(t *testing.T) {
	subtitles := []struct {
		startCut  string
		finishCut string
		sample    string
		truncated string
	}{
		{"0s", "5s", "testdata/sample.srt", "testdata/sample.srt"},
		{"0s", "3s", "testdata/sample.srt", "testdata/truncated_sample_a.srt"},
		{"0s", "4s", "testdata/sample.srt", "testdata/truncated_sample_b.srt"},
		{"1s500ms", "4s", "testdata/sample.srt", "testdata/truncated_sample_c.srt"},
		{"1s500ms", "6s", "testdata/sample.srt", "testdata/truncated_sample_d.srt"},
	}

	for _, subtitle := range subtitles {
		truncatedSubtitle, _ := Load(subtitle.truncated)
		startAt, _ := time.ParseDuration(subtitle.startCut)
		finishAt, _ := time.ParseDuration(subtitle.finishCut)

		sub, _ := Load(subtitle.sample)
		sub.Cut(startAt, finishAt)

		assert.EqualValues(t, truncatedSubtitle, sub)
	}
}
