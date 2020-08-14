package subtitles

import (
	"subtitles/utils"
)

// Subtitle struct
type Subtitle struct{}

// LoadFile method
func LoadFile(file string) (Subtitle, error) {
	if utils.FileExists(file) == false {
		return Subtitle{}, &ErrInvalidFile{file: file}
	}

	return Subtitle{}, nil
}
