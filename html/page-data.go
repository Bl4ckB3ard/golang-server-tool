package html

import (
	"path/filepath"

	"github.com/Bl4ckB3ard/golang-server-tool/dirparser"
)

type PageItem struct {
	IsDir        bool
	Name         string
	Size         string
	ViewHref     string
	DownloadHref string
	RelativePath string
}

type PageData struct {
	RootDirectory  string
	Items          []PageItem
	LenOfDirectory int
}

func GetPageData(directory string, rootFS dirparser.RootFS) PageData {
	var pd PageData
	items := dirparser.GetDirItems(directory, rootFS)

	pageItems := make([]PageItem, len(items))

	for idx, item := range items {
		pi := PageItem{
			Name:         item.BaseName,
			IsDir:        item.IsDir,
			Size:         item.Size,
			ViewHref:     filepath.Clean(item.RelativePath) + "?view=true",
			DownloadHref: filepath.Clean(item.RelativePath) + "?download=true",
			RelativePath: item.RelativePath,
		}

		pageItems[idx] = pi
	}

	pd = PageData{
		RootDirectory:  directory,
		Items:          pageItems,
		LenOfDirectory: len(pageItems),
	}
	return pd
}
