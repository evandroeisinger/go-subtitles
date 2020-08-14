package subtitles

import (
	"subtitles/utils"
)

// Block struct
type Block struct {
	lines   []string
	startAt int
	endAt   int
}

// Subtitle struct
type Subtitle struct {
	content string
	blocks  []Block
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

	blocks, err := ParserForFile(path).parse(&content)

	subtitle := Subtitle{
		content: content,
		blocks:  blocks,
	}

	return subtitle, err
}
