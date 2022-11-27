package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func IsValidFile(file string) bool {
	fullPath, err := filepath.Abs(file)
	if err != nil {
		return false
	}
	fileHandle, err := os.Open(fullPath)
	defer fileHandle.Close()
	if err != nil {
		return false
	}
	info, err := fileHandle.Stat()
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func IsValidDirectory(dir string) bool {
	fullPath, err := filepath.Abs(dir)
	if err != nil {
		return false
	}
	fileHandle, err := os.Open(fullPath)
	defer fileHandle.Close()
	if err != nil {
		return false
	}
	info, err := fileHandle.Stat()
	if err != nil {
		return false
	}
	return info.IsDir()
}

func EscapePath(path string) string {
	var escaped string
	escaped = strings.Replace(path, "\\", "\\\\", -1)
	escaped = strings.Replace(escaped, "[", "\\[", -1)
	escaped = strings.Replace(escaped, "]", "\\]", -1)

	escaped = strings.Replace(escaped, "(", "\\(", -1)
	escaped = strings.Replace(escaped, ")", "\\)", -1)

	escaped = strings.Replace(escaped, "{", "\\{", -1)
	escaped = strings.Replace(escaped, "}", "\\}", -1)

	escaped = strings.Replace(escaped, "*", "\\*", -1)
	escaped = strings.Replace(escaped, "+", "\\+", -1)
	escaped = strings.Replace(escaped, "?", "\\?", -1)
	escaped = strings.Replace(escaped, "|", "\\|", -1)
	escaped = strings.Replace(escaped, "^", "\\^", -1)
	escaped = strings.Replace(escaped, "$", "\\$", -1)
	escaped = strings.Replace(escaped, ".", "\\.", -1)

	return escaped
}
