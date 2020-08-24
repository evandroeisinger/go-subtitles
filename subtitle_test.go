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
