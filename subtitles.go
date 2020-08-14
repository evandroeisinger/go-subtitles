package subtitles

import (
	"os"
)

// Subtitle struct
type Subtitle struct{}

// LoadFile method
func LoadFile(file string) (Subtitle, error) {
	if _, err := os.Stat(file); err != nil {
		return Subtitle{}, &ErrInvalidFile{file: file}
	}

	return Subtitle{}, nil
}
