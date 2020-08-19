package subtitles

import "io"

// SRT subtitle format
type SRT struct {
}

// SRTExtension extension for format
const SRTExtension string = ".srt"

// NewSRT returns SRT format instance
func NewSRT() *SRT {
	return &SRT{}
}

// Parse from SRT subtitle format
func (f *SRT) Parse(r io.Reader) (s *Subtitle, err error) {
	s = &Subtitle{}
	return s, err
}

// Write to SRT subtitle format
func (f *SRT) Write(w io.Writer) (err error) {
	return err
}
