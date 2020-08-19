package subtitles

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEmptySubtitle(t *testing.T) {
	subtitle, err := NewSRT().Parse(strings.NewReader(""))

	assert.Equal(t, &Subtitle{}, subtitle)
	assert.Nil(t, err)
}

func TestReadSimpleBlock(t *testing.T) {
	content := strings.NewReader("1\n00:00:00,000 --> 00:00:01,000\nLorem ipsum dolor sit amet\n")

	subtitle, err := NewSRT().Parse(content)

	assert.Equal(t, 1, len(subtitle.blocks))
	assert.Nil(t, err)
}
