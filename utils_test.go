package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripBom(t *testing.T) {
	assert.Equal(t, "Lorem ipsum", StripBOM("\ufeffLorem ipsum"))
}
