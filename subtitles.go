package subtitles

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

// Load method
func Load(path string) (Subtitle, error) {
	subtitle := Subtitle{
		content: "",
		blocks:  []Block{},
	}

	return subtitle, nil
}
