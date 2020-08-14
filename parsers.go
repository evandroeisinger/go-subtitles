package subtitles

// Parser interface
type Parser interface {
	parse(content *string) ([]Block, error)
}

// ParserForFile method
func ParserForFile(path string) Parser {
	// for test pourpuse it returns SRTParser
	return &SRTParser{}
}
