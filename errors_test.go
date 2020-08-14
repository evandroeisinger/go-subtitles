package subtitles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidFile(t *testing.T) {
	err := &ErrInvalidFile{file: "invalid-file.srt"}

	assert.EqualError(t, err, "Invalid file: invalid-file.srt")
}

func TestErrInvalidFileContent(t *testing.T) {
	err := &ErrInvalidFileContent{
		file:    "invalid-file-content",
		content: "Invalid Content",
	}

	assert.EqualError(t, err, fmt.Sprintf("Invalid file content: %s\n%s", err.file, err.content))
}
