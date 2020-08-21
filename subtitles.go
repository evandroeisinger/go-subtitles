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
func ParserForFile(f string) (Parser, error) {
	fileExtension := filepath.Ext(f)

	var parser Parser
	var err error

	switch fileExtension {
	case SRTExtension:
		parser = NewSRTParser()
	default:
		err = &ErrInvalidFile{f, "Parser for extension not found"}
	}

	return parser, err
}

// FormatterForFile returns parser for subtitle format
func FormatterForFile(f string) (Formatter, error) {
	fileExtension := filepath.Ext(f)

	var formatter Formatter
	var err error

	switch fileExtension {
	case SRTExtension:
		formatter = NewSRTFormatter()
	default:
		err = &ErrInvalidFile{f, "Formatter for extension not found"}
	}

	return formatter, err
}

// Load subtitle
func Load(p string) (*Subtitle, error) {
	file, err := OpenFile(p)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	parser, err := ParserForFile(p)
	if err != nil {
		return nil, err
	}

	subtitle, err := parser.Parse(file)
	return subtitle, err
}

// Write subtitle
func Write(s *Subtitle, p string) (int, error) {
	parser, err := FormatterForFile(p)
	if err != nil {
		return 0, err
	}

	content, err := parser.Format(s)
	if err != nil {
		return 0, err
	}

	file, err := CreateFile(p)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	n, err := file.WriteString(content)

	return n, err
}
