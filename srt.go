package subtitles

import "io"

// SRTExtension extension for format
const SRTExtension string = ".srt"

// SRTParser subtitle format
type SRTParser struct{}

// NewSRTParser returns SRT format instance
func NewSRTParser() *SRTParser {
	return &SRTParser{}
}

// Parse SRT subtitle content
func (p *SRTParser) Parse(r io.Reader) (s *Subtitle, err error) {
	s = NewSubtitle()

	return s, err
}
