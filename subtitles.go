package subtitles

import (
	"sort"
	"time"
)

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

// Concat subtitles
func Concat(subs ...*Subtitle) (sub *Subtitle, err error) {
	blocks := []*Block{}
	durationCorrection := time.Second * 0

	for _, sub := range subs {
		for _, b := range sub.Blocks {
			block := NewBlock()
			block.Lines = b.Lines
			block.StartAt = b.StartAt + durationCorrection
			block.FinishAt = b.FinishAt + durationCorrection

			blocks = append(blocks, block)
		}

		lastBlockIndex := len(sub.Blocks) - 1
		durationCorrection = sub.Blocks[lastBlockIndex].FinishAt
	}

	return NewSubtitle(blocks...), err
}

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
