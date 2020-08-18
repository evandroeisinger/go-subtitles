package subtitles

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEmptySubtitle(t *testing.T) {
	subtitle, err := NewSRT().Read(strings.NewReader(""))

	assert.Equal(t, &Subtitle{}, subtitle)
	assert.Nil(t, err)
}
