package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSRTParse(t *testing.T) {
	content := ""
	parser := SRTParser{}
	blocks, err := parser.parse(&content)

	assert.Equal(t, blocks, make([]Block, 5))
	assert.Nil(t, err)
}
