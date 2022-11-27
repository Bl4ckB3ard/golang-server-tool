package dirparser

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
