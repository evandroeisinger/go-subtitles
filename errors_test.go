package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidFile(t *testing.T) {
	err := &ErrInvalidFile{file: "invalid-file.srt"}

	assert.EqualError(t, err, "Invalid file: invalid-file.srt")
}
