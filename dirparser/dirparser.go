package dirparser

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

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

func GetRootDirPaths(dir string) ([]string, error) {
	var names []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("ERROR: getting root directory's content")
			os.Exit(2)
		}

		names = append(names, path+"/"+d.Name())
		return nil
	})

	if err != nil {
		return []string{}, err
	}
	return names, nil
}
