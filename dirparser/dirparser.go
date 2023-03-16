package dirparser

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Item struct {
	BaseName     string
	FullPath     string
	Size         string
	IsDir        bool
	IsViewAble   bool
	RelativePath string
}

type RootFS struct {
	RootPath string
	Items    []Item
}

func ParseFileSize(size int64) string {
	bytee := size < 1000
	KB := size >= 1000 && size < 1000000
	MB := size >= 1000000 && size < 1000000000
	GB := size >= 1000000000

	switch {
	case bytee:
		return strings.Replace(fmt.Sprintf("%d Bytes", size), ".00", "", 1)
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
	return fmt.Sprintf("%d Items", len(items))
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


func GetDirItems(fullDirPath string, rootFS RootFS) []Item {
	var items []Item
	for _, item := range rootFS.Items {
		sameName := filepath.Clean(item.FullPath) == filepath.Clean(fullDirPath)

		escapedPath := EscapePath(fullDirPath)

		re := regexp.MustCompile("^(" + escapedPath + ")" + "\\/[^/]+\\/?$")

		match := re.MatchString(filepath.Clean(item.FullPath))

		if match && !sameName {
			items = append(items, item)
		}
	}

	return items
}


func GetRootFS(rootDir string) (RootFS, error) {
	FS := RootFS{RootPath: rootDir}
	var items []Item

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("ERROR: getting root directory's content")
			os.Exit(2)
		}

		info, _ := d.Info()

		var fullPath string
		var size string
		if !info.IsDir() {
			fullPath = path
			size = ParseFileSize(info.Size())
		} else {
			fullPath = path + "/"
			size = GetDirSize(fullPath)
		}

		items = append(items, Item{
			BaseName:     d.Name(),
			FullPath:     fullPath,
			RelativePath: strings.Replace(fullPath, filepath.Clean(rootDir), "", 1),
			IsDir:        d.IsDir(),
			Size:         size,
		})

		return nil
	})

	if err != nil {
		return RootFS{}, err
	}

	FS.Items = items

	return FS, nil
}
