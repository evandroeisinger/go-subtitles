package subtitles

import "strings"

// BOMUnicode value
const BOMUnicode = "\ufeff"

// StripBOM removes bytes order mark from string
func StripBOM(s string) string {
	return strings.Trim(s, BOMUnicode)
}
