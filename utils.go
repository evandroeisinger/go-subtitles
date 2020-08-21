package subtitles

import (
	"os"
	"path/filepath"
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

// PathExist checks if its a valid path
func PathExist(p string) bool {
	if _, err := os.Stat(p); err != nil {
		return false
	}

	return true
}

// OpenFile validate and return File reader
func OpenFile(p string) (*os.File, error) {
	if !PathExist(p) {
		return nil, &ErrInvalidFile{p, "File not exist"}
	}

	if !IsFile(p) {
		return nil, &ErrInvalidFile{p, "Its not a file"}
	}

	if IsEmptyFile(p) {
		return nil, &ErrInvalidFile{p, "Empty file"}
	}

	file, err := os.Open(p)

	return file, err
}

// CreateFile validate and return File reader
func CreateFile(p string) (*os.File, error) {
	if !PathExist(filepath.Dir(p)) {
		return nil, &ErrInvalidFile{p, "File path not exist"}
	}

	if PathExist(p) {
		return nil, &ErrInvalidFile{p, "File already exist"}
	}

	file, err := os.Create(p)
	if err != nil {
		return nil, err
	}

	return file, nil
}
