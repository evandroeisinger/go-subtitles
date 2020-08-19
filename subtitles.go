package subtitles

import (
	"fmt"
	"io"
	"path/filepath"
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

// NewSubtitle returns subtitle instance
func NewSubtitle() *Subtitle {
	return &Subtitle{}
}

// ErrUnsupportedExtension error
type ErrUnsupportedExtension struct {
	extension string
}

func (e *ErrUnsupportedExtension) Error() string {
	return fmt.Sprintf("Unsupported extension: %s", e.extension)
}

// Parser interface
type Parser interface {
	Parse(r io.Reader) (*Subtitle, error)
}

// ParserForFile returns parser for subtitle format
func ParserForFile(f string) (p Parser, err error) {
	fileExtension := filepath.Ext(f)

	switch fileExtension {
	case SRTExtension:
		p = NewSRTParser()
	default:
		err = &ErrUnsupportedExtension{
			extension: fileExtension,
		}
	}

	return p, err
}

// Load method
func Load(path string) (s Subtitle, err error) {
	s = Subtitle{
		content: "",
		blocks:  []Block{},
	}

	return s, err
}
