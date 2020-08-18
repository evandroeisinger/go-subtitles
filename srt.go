package subtitles

import "io"

// SRT subtitle format
type SRT struct {
}

// NewSRT returns SRT format instance
func NewSRT() *SRT {
	return &SRT{}
}

// Read from SRT subtitle format
func (format *SRT) Read(reader io.Reader) (*Subtitle, error) {
	return &Subtitle{}, nil
}

// Write to SRT subtitle format
// func (format *SRT) Write(writer io.Writer) error {
// 	return nil
// }
