package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func IsValidPort(port string) bool {
	num, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return false
	}
	return num >= 1 && num < 65536
}

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

func CheckServerErr(err error) {
	if err != nil {
		fmt.Println("ERROR starting server")
		os.Exit(1)
	}
}

// TODO:
func IsViewAble(fullPath string) bool {
	return true
}
