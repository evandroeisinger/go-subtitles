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
