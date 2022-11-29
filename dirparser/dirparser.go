package dirparser

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

func GetDirItems(fullDirPath string, rootFS RootFS) []Item {
	var items []Item
	for _, item := range rootFS.Items {
		sameName := filepath.Clean(item.FullPath) == filepath.Clean(fullDirPath)

		escapedPath := utils.EscapePath(fullDirPath)

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
			size = utils.ParseFileSize(info.Size())
		} else {
			fullPath = path + "/"
			size = utils.GetDirSize(fullPath)
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
