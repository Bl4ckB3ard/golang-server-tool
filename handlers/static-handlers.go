package handlers

import (
	"bytes"
	"net/http"
	"time"

	"github.com/Bl4ckB3ard/golang-server-tool/static"
)

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "favicon.ico", time.Now(), bytes.NewReader(static.FaviconIcoFile))
	return
}

func FileLight(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "file_icon_light.png", time.Now(), bytes.NewReader(static.FileIconLight))
	return
}

func FileDark(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "file_icon_dark.png", time.Now(), bytes.NewReader(static.FileIconDark))
}

func FolderLight(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "file_icon_light.png", time.Now(), bytes.NewReader(static.FolderIconLight))
}

func FolderDark(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "folder_icon_dark.png", time.Now(), bytes.NewReader(static.FolderIconDark))
}

func LogoLight(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "logo_light.png", time.Now(), bytes.NewReader(static.LogoLight))
}

func LogoDark(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "logo_dark.png", time.Now(), bytes.NewReader(static.LogoDark))
}

func ArrowLight(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "sort_arrow_icon_light.png", time.Now(), bytes.NewReader(static.SortArrowLight))
}

func ArrowDark(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, "sort_arrow_icon_dark.png", time.Now(), bytes.NewReader(static.SortArrowDark))
}
