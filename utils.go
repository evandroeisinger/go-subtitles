package subtitles

import (
	"os"
	"strings"
)

// BOMUnicode value
const BOMUnicode = "\ufeff"

// StripBOM removes bytes order mark from string
func StripBOM(s string) string {
	return strings.Trim(s, BOMUnicode)
}

// IsEmptyFile checks if file has content
func IsEmptyFile(p string) bool {
	fileInfo, err := os.Stat(p)

	if err != nil {
		return true
	}

	return fileInfo.Size() == 0
}
