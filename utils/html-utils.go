package utils

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// TODO:
func IsViewAble(fullPath string) bool {
	// validExts := []string{".mp4", ".mp3", ".json", ".html", ".css", ".js", ".py", ".go", ".c", ".jpg", ".jpeg", ".png", ".ogg", ".webm", ""}

	return true
}

func ParseFileSize(size int64) string {
	bytee := size < 1000
	KB := size >= 1000 && size < 1000000
	MB := size >= 1000000 && size < 1000000000
	GB := size >= 1000000000

	switch {
	case bytee:
		return fmt.Sprintf("%d Bytes", size)
	case KB:
		return strings.Replace(fmt.Sprintf("%.2f KB", math.Round(float64(size/1000))), ".00", "", 1)
	case MB:
		return strings.Replace(fmt.Sprintf("%.2f MB", math.Round(float64(size/1000000))), ".00", "", 1)
	case GB:
		return strings.Replace(fmt.Sprintf("%.2f GB", math.Round(float64(size/1000000000))), ".00", "", 1)
	default:
		return "0"
	}
}

func GetDirSize(p string) string {
	handle, _ := os.Open(p)
	defer handle.Close()
	items, _ := handle.ReadDir(0)
	return fmt.Sprintf("%d, items", len(items))
}
