package subtitles

import (
	"fmt"
	"io"
	"path/filepath"
	"time"
)

// Block struct
type Block struct {
	lines    []string
	startAt  time.Duration
	finishAt time.Duration
}

// NewBlock returns block instance
func NewBlock() *Block {
	return &Block{}
}

// Subtitle struct
type Subtitle struct {
	blocks []*Block
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

// ParserForFile returns parser for subtitle format
func ParserForFile(f string) (p Parser, err error) {
	fileExtension := filepath.Ext(f)

	switch fileExtension {
	case SRTExtension:
		p = NewSRTParser()
	default:
		err = &ErrInvalidFile{f, "Extension not supported"}
	}

	return p, err
}

// Load method
func Load(p string) (*Subtitle, error) {
	file, err := OpenFile(p)
	if err != nil {
		return nil, err
	}

	parser, err := ParserForFile(p)
	if err != nil {
		return nil, err
	}

	subtitle, err := parser.Parse(file)
	return subtitle, err
}
