package subtitles

import (
	"fmt"
	"io"
	"path/filepath"
	"time"
)

// Block struct
type Block struct {
	Lines    []string      `json:"lines"`
	StartAt  time.Duration `json:"startAt"`
	FinishAt time.Duration `json:"finishAt"`
}

// NewBlock returns block instance
func NewBlock() *Block {
	return &Block{}
}

// Subtitle struct
type Subtitle struct {
	Blocks []*Block `json:"blocks"`
}

// NewSubtitle returns subtitle instance
func NewSubtitle() *Subtitle {
	return &Subtitle{}
}

// ErrInvalidSubtitle error
type ErrInvalidSubtitle struct {
	format string
	reason string
}

func (e *ErrInvalidSubtitle) Error() string {
	return fmt.Sprintf("Invalid subtitle %s: %s", e.format, e.reason)
}

// ErrInvalidFile error
type ErrInvalidFile struct {
	file   string
	reason string
}

func (e *ErrInvalidFile) Error() string {
	return fmt.Sprintf("Invalid file %s: %s", e.file, e.reason)
}

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

// Load subtitle
func Load(path string) (sub *Subtitle, err error) {
	file, err := OpenFile(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	parser, err := ParserForFile(path)
	if err != nil {
		return nil, err
	}

	sub, err = parser.Parse(file)
	return sub, err
}

// Write subtitle
func Write(s *Subtitle, path string) (n int, err error) {
	parser, err := FormatterForFile(path)
	if err != nil {
		return 0, err
	}

	content, err := parser.Format(s)
	if err != nil {
		return 0, err
	}

	file, err := CreateFile(path)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	n, err = file.WriteString(content)

	return n, err
}
