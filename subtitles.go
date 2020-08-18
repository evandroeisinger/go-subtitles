package subtitles

import (
	"fmt"
	"io"
	"path/filepath"
)

// Format interface
type Format interface {
	Read(reader io.Reader) (*Subtitle, error)
	Write(writer io.Writer) error
}

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

// ErrUnsupportedExtension error
type ErrUnsupportedExtension struct {
	extension string
}

func (e *ErrUnsupportedExtension) Error() string {
	return fmt.Sprintf("Unsupported extension: %s", e.extension)
}

// FormatForFile returns format
func FormatForFile(path string) (Format, error) {
	fileExtension := filepath.Ext(path)

	var format Format
	var err error

	switch fileExtension {
	case ".srt":
		format = NewSRT()
	default:
		err = &ErrUnsupportedExtension{
			extension: fileExtension,
		}
	}

	return format, err
}

// Load method
func Load(path string) (Subtitle, error) {
	subtitle := Subtitle{
		content: "",
		blocks:  []Block{},
	}

	return subtitle, nil
}
