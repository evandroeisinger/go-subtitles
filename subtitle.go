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

// BlockSorter sorts blocks by startAt
type BlockSorter []*Block

func (b BlockSorter) Len() int           { return len(b) }
func (b BlockSorter) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BlockSorter) Less(i, j int) bool { return b[i].StartAt < b[j].StartAt }
