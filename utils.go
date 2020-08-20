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

// IsFile checks if its a file
func IsFile(p string) bool {
	fileInfo, err := os.Stat(p)

	if err != nil {
		return false
	}

	return !fileInfo.IsDir()
}

// IsEmptyFile checks if file has content
func IsEmptyFile(p string) bool {
	fileInfo, err := os.Stat(p)

	if err != nil {
		return true
	}

	return fileInfo.Size() == 0
}

// FileExist checks if its a valid file
func FileExist(p string) bool {
	if _, err := os.Stat(p); err != nil {
		return false
	}

	return true
}
