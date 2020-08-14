package subtitles

import (
	"fmt"
)

// ErrInvalidFile error
type ErrInvalidFile struct {
	file string
}

func (e *ErrInvalidFile) Error() string {
	return fmt.Sprintf("Invalid file: %s", e.file)
}

// ErrInvalidFileContent error
type ErrInvalidFileContent struct {
	file    string
	content string
}

func (e *ErrInvalidFileContent) Error() string {
	return fmt.Sprintf("Invalid file content: %s\n%s", e.file, e.content)
}
