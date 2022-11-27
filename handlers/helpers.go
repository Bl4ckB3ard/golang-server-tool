package handlers

import (
	"path/filepath"

	"github.com/Bl4ckB3ard/golang-server-tool/dirparser"
)

func IsInRoot(path string, FS dirparser.RootFS) (bool, dirparser.Item) {
	for _, item := range FS.Items {
		if filepath.Clean(path) == filepath.Clean(item.RelativePath) {
			return true, item
		}
	}
	return false, dirparser.Item{}
}
