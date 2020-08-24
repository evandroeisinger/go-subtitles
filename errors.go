package subtitles

import "fmt"

// ErrInvalidSubtitle error
type ErrInvalidSubtitle struct {
	format string
	reason string
}

func (e *ErrInvalidSubtitle) Error() string {
	return fmt.Sprintf("Invalid subtitle %s: %s", e.format, e.reason)
}

// ErrInvalidFile error
type ErrInvalidFile struct {
	file   string
	reason string
}

func (e *ErrInvalidFile) Error() string {
	return fmt.Sprintf("Invalid file %s: %s", e.file, e.reason)
}
