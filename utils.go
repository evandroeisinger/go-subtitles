package subtitles

import "strings"

// StripBOM removes bytes order mark from string
func StripBOM(s string) string {
	return strings.TrimPrefix(s, "\ufeff")
}
