package dirparser

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

type Item struct {
	BaseName   string
	FullPath   string
	IsDir      bool
	IsViewAble bool
}

func GetDirNames(dir string) ([]string, error) {
	fileHandle, err := os.Open(dir)
	if !utils.IsValidDirectory(dir) {
		return []string{}, errors.New(fmt.Sprintf("%v is not a valid directory", dir))
	}

	content, err := fileHandle.ReadDir(0)
	if err != nil {
		return []string{}, err
	}

	names := make([]string, len(content))

	for idx, val := range content {
		names[idx] = val.Name()
	}

	return names, nil
}

func GetRootFS(dir string) ([]Item, error) {
	var FS []Item

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("ERROR: getting root directory's content")
			os.Exit(2)
		}

		fullPath := path + "/" + d.Name()

		FS = append(FS, Item{BaseName: d.Name(), FullPath: fullPath, IsDir: d.IsDir(), IsViewAble: utils.IsViewAble(fullPath)})
		return nil
	})

	if err != nil {
		return []Item{}, err
	}

	return FS, nil
}
