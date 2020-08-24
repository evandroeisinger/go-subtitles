package subtitles

import (
	"io"
	"path/filepath"
)

// Parser interface
type Parser interface {
	Parse(r io.Reader) (*Subtitle, error)
}

// Formatter interface
type Formatter interface {
	Format(s *Subtitle) (string, error)
}

// ParserForFile returns parser for subtitle format
func ParserForFile(path string) (p Parser, err error) {
	fileExtension := filepath.Ext(path)

	switch fileExtension {
	case SRTExtension:
		p = NewSRTParser()
	default:
		err = &ErrInvalidFile{path, "Parser for extension not found"}
	}

	return p, err
}

// FormatterForFile returns parser for subtitle format
func FormatterForFile(path string) (f Formatter, err error) {
	fileExtension := filepath.Ext(path)

	switch fileExtension {
	case SRTExtension:
		f = NewSRTFormatter()
	default:
		err = &ErrInvalidFile{path, "Formatter for extension not found"}
	}

	return f, err
}
