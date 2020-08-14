package subtitles

import (
	"subtitles/utils"
)

// Subtitle struct
type Subtitle struct{}

// LoadFromFile method
func LoadFromFile(path string) (Subtitle, error) {
	if utils.FileExists(path) == false {
		return Subtitle{}, &ErrInvalidFile{file: path}
	}

	return Subtitle{}, nil
}
