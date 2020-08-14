package subtitles

import (
	"subtitles/utils"
)

// Subtitle struct
type Subtitle struct {
	content string
}

// LoadFromFile method
func LoadFromFile(path string) (Subtitle, error) {
	if utils.FileExists(path) == false {
		return Subtitle{}, &ErrInvalidFile{file: path}
	}

	content := utils.LoadFileContent(path)
	if len(content) == 0 {
		return Subtitle{}, &ErrInvalidFileContent{
			file:    path,
			content: content,
		}
	}

	subtitle := Subtitle{
		content: content,
	}

	return subtitle, nil
}
