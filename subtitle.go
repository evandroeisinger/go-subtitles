package subtitles

import (
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
func (sub *Subtitle) Shift(d time.Duration) *Subtitle {
	for _, block := range sub.Blocks {
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

	return sub
}

// Cut subtitle
func (sub *Subtitle) Cut(s time.Duration, f time.Duration) *Subtitle {
	blocks := []*Block{}
	for _, block := range sub.Blocks {
		// break if starts after cut duration
		if block.StartAt >= f {
			break
		}

		// skip if starts and finishes before cut duration
		if block.StartAt <= s && block.FinishAt <= s {
			continue
		}

		// fix duration if block ends after cut
		if block.FinishAt > f {
			block.FinishAt = f
		}

		// fix duration if block starts before cut
		if block.StartAt < s {
			block.StartAt = s
		}

		blocks = append(blocks, block)
	}

	sub.Blocks = blocks
	return sub
}

// NewSubtitle returns subtitle instance
func NewSubtitle(blocks ...*Block) *Subtitle {
	return &Subtitle{Blocks: blocks}
}

// BlockSorter sorts blocks by startAt
type BlockSorter []*Block

func (b BlockSorter) Len() int           { return len(b) }
func (b BlockSorter) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BlockSorter) Less(i, j int) bool { return b[i].StartAt < b[j].StartAt }
