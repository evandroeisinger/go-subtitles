package subtitles

// SRTParser type
type SRTParser struct{}

func (p *SRTParser) parse(content *string) ([]Block, error) {
	// for test pourpose, return 5 empty blocks
	return make([]Block, 5), nil
}
