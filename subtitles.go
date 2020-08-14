package subtitles

import (
	"subtitles/utils"
)

// Subtitle struct
type Subtitle struct{}

// LoadFile method
func LoadFile(path string) (Subtitle, error) {
	if utils.FileExists(path) == false {
		return Subtitle{}, &ErrInvalidFile{file: path}
	}

	return Subtitle{}, nil
}
