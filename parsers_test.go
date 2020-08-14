package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ParserForFile method
func TestParserForFile(t *testing.T) {
	filePaths := []struct {
		path   string
		parser Parser
	}{
		{path: "testdata/sample.srt", parser: &SRTParser{}},
	}

	for _, file := range filePaths {
		assert.Equal(t, ParserForFile(file.path), file.parser)
	}
}
