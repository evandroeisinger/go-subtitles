package subtitles

import (
	"fmt"
	"io"
	"path/filepath"
	"sort"
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

// Shift subtitle
func (s *Subtitle) Shift(d time.Duration) *Subtitle {
	for _, block := range s.Blocks {
		startAt := block.StartAt + d
		if startAt < 0 {
			startAt, _ = time.ParseDuration("0s")
		}

		finishAt := block.FinishAt + d
		if finishAt < 0 {
			finishAt, _ = time.ParseDuration("0s")
		}

		block.StartAt = startAt
		block.FinishAt = finishAt
	}

	return s
}

// NewSubtitle returns subtitle instance
func NewSubtitle(blocks ...*Block) *Subtitle {
	return &Subtitle{Blocks: blocks}
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
func Write(s *Subtitle, path string) (c string, err error) {
	parser, err := FormatterForFile(path)
	if err != nil {
		return "", err
	}

	c, err = parser.Format(s)
	if err != nil {
		return "", err
	}

	file, err := CreateFile(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = file.WriteString(c)
	return c, err
}

// BlockSorter sorts blocks by startAt
type BlockSorter []*Block

func (b BlockSorter) Len() int           { return len(b) }
func (b BlockSorter) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BlockSorter) Less(i, j int) bool { return b[i].StartAt < b[j].StartAt }

// Merge subtitles
func Merge(subs ...*Subtitle) (sub *Subtitle, err error) {
	blocks := []*Block{}

	for _, s := range subs {
		blocks = append(blocks, s.Blocks...)
	}

	// Sort merged blocks
	sort.Sort(BlockSorter(blocks))

	return NewSubtitle(blocks...), err
}
