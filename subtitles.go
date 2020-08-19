package subtitles

import (
	"fmt"
	"io"
	"path/filepath"
)

// Format interface
type Format interface {
	Parse(r io.Reader) (*Subtitle, error)
	Write(w io.Writer) error
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

// FormatForFile returns format
func FormatForFile(p string) (f Format, err error) {
	fileExtension := filepath.Ext(p)

	switch fileExtension {
	case SRTExtension:
		f = NewSRT()
	default:
		err = &ErrUnsupportedExtension{
			extension: fileExtension,
		}
	}

	return f, err
}

// Load method
func Load(path string) (s Subtitle, err error) {
	s = Subtitle{
		content: "",
		blocks:  []Block{},
	}

	return s, err
}
