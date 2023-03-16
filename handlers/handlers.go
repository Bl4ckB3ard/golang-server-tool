package handlers

import (
	"bytes"
	"encoding/base64"
	"compress/gzip"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Bl4ckB3ard/golang-server-tool/config"
	"github.com/Bl4ckB3ard/golang-server-tool/dirparser"
	"github.com/Bl4ckB3ard/golang-server-tool/page"
	"github.com/Bl4ckB3ard/golang-server-tool/static"
)

func CheckMethod(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Method %s is not allowed.\n", r.Method), http.StatusMethodNotAllowed)
		return errors.New("Invalid Request Method")
	}
	return nil
}


func DefaultHeaders(w http.ResponseWriter) {
	w.Header().Add("Server", "GoLang")
}


func acceptsGzip(r *http.Request) bool {
    for _, i := range strings.Split(r.Header.Get("Accepts-Encoding"), ", ") {
        if i == "gzip" {
            return true
        } 
    }

    return false
}


func gunzip(zipedData []byte) []byte {
    gr, _ := gzip.NewReader(bytes.NewReader(zipedData))

    gunziped, _ := io.ReadAll(gr)
    return gunziped
}


func IsInRoot(path string, FS dirparser.RootFS) (bool, dirparser.Item) {
	for _, item := range FS.Items {
		if filepath.Clean(path) == filepath.Clean(item.RelativePath) {
			return true, item
		}
	}
	return false, dirparser.Item{}
}


func StaticHandler(w http.ResponseWriter, r *http.Request) {
    CheckMethod(w, r)

    var serveFile []byte
    var fName string

    switch (filepath.Base(r.URL.Path)) {
    case "file_icon_light.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.FileIconLight)
        fName = "file_icon_light.png"
        break
    case "folder_icon_light.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.FolderIconLight)
        fName = "folder_icon_light.png"
        break
    case "sort_arrow_icon_light.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.SortArrowLight)
        fName = "sort_arrow_icon_light.png"
        break
    case "logo_light.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.LogoLight)
        fName = "logo_light.png"
        break

    case "file_icon_dark.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.FileIconDark)
        fName = "file_icon_dark.png"
        break
    case "folder_icon_dark.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.FolderIconDark)
        fName = "folder_icon_dark.png"
        break
    case "sort_arrow_icon_dark.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.SortArrowDark)
        fName = "sort_arrow_icon_dark.png"
        break
    case "logo_dark.png":
        serveFile, _ = base64.StdEncoding.DecodeString(static.LogoDark)
        fName = "logo_dark.png"
        break
    default:
        http.Error(w, fmt.Sprintf("404 not found %v", filepath.Base(r.URL.Path)), http.StatusNotFound)
        return
    }

    if !acceptsGzip(r) {
        http.ServeContent(w, r, fName, time.Now(), bytes.NewReader(gunzip(serveFile)))
        return
    }

    w.Header().Set("Content-Encoding", "gzip")
    http.ServeContent(w, r, fName, time.Now(), bytes.NewReader(serveFile))
    return
}


func FileHandler(w http.ResponseWriter, r *http.Request) {
	if err := CheckMethod(w, r); err != nil {
		return
	}

	content, _ := os.Open(config.ARGS.FilePath)

	DefaultHeaders(w)

	http.ServeContent(w, r, filepath.Base(config.ARGS.FilePath), time.Now(), content)
}


func handleDirectory(w http.ResponseWriter, i dirparser.Item, r *http.Request) {
	if err := CheckMethod(w, r); err != nil {
		return
	}

	if !i.IsDir {
		http.Error(w, fmt.Sprintf("%v is a file view or download parameter is mandatory", i.BaseName), http.StatusBadRequest)
		return
	}

	pageData := page.GetPageData(filepath.Clean(i.FullPath), config.RootFS)
	var err error
	var tmpl *template.Template
	if config.ARGS.Theme == "light" {
		tmpl, err = template.New(i.BaseName).Parse(static.LightThemeTemplate)
	} else {
		tmpl, err = template.New(i.BaseName).Parse(static.DarkThemeTemplate)
	}

	if err != nil {
		http.Error(w, "ERROR: invalid template", http.StatusInternalServerError)
		fmt.Printf("%v\nERROR: couldn't parse template\n", err)
		os.Exit(2)
	}

	DefaultHeaders(w)
	tmpl.Execute(w, pageData)
	return
}


func handleRoot(w http.ResponseWriter, r *http.Request) {
	if err := CheckMethod(w, r); err != nil {
		return
	}

	pageData := page.GetPageData(config.ARGS.DirectoryPath, config.RootFS)
	var err error
	var tmpl *template.Template
	if config.ARGS.Theme == "light" {
		tmpl, err = template.New("home").Parse(static.LightThemeTemplate)
	} else {
		tmpl, err = template.New("home").Parse(static.DarkThemeTemplate)
	}

	if err != nil {
		http.Error(w, "ERROR: invalid template", http.StatusInternalServerError)
		fmt.Printf("%v\nERROR: couldn't parse template\n", err)
		os.Exit(2)
	}

	DefaultHeaders(w)

	tmpl.Execute(w, pageData)

	return
}


func handleView(w http.ResponseWriter, i dirparser.Item, r *http.Request) {
	if err := CheckMethod(w, r); err != nil {
		return
	}

	if i.IsDir {
		http.Error(w, "Can not use view and download parameters on directory", http.StatusBadRequest)
		return
	}
	content, _ := os.Open(i.FullPath)
	w.Header().Add("Content-Disposition", "inline")
	http.ServeContent(w, r, i.BaseName, time.Now(), content)
	return
}


func handleDownload(w http.ResponseWriter, i dirparser.Item, r *http.Request) {
	if err := CheckMethod(w, r); err != nil {
		return
	}

	if i.IsDir {
		http.Error(w, "Cannont use view and download parameters on directory", http.StatusBadRequest)
		return
	}
	content, _ := os.Open(i.FullPath)
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", i.BaseName))
	http.ServeContent(w, r, i.BaseName, time.Now(), content)
	return
}


func MainHandler(w http.ResponseWriter, r *http.Request) {
	if err := CheckMethod(w, r); err != nil {
		return
	}

	if r.URL.EscapedPath() == "/" {
		handleRoot(w, r)
		return
	} else {
		path := r.URL.Path
		view, err := strconv.ParseBool(r.URL.Query().Get("view"))
		if err != nil && r.URL.Query().Get("view") != "" {
			http.Error(w, fmt.Sprintf("INVALID GET parameter \"%v\": view paramter is a bool", r.URL.Query().Get("view")), http.StatusBadRequest)
			return
		}

		download, err := strconv.ParseBool(r.URL.Query().Get("download"))
		if err != nil && r.URL.Query().Get("download") != "" {
			http.Error(w, fmt.Sprintf("INVALID GET parameter \"%v\": view paramter is a bool", r.URL.Query().Get("download")), http.StatusBadRequest)
			return
		}

		viewOnly := view && !download
		downloadOnly := download && !view
		both := view && download
		none := !view && !download

		valid, item := IsInRoot(path, config.RootFS)
		if !valid {
			http.Error(w, fmt.Sprintf("404 not found %v", path), http.StatusNotFound)
			return
		}

		switch {
		case viewOnly:
			handleView(w, item, r)
			return
		case downloadOnly:
			handleDownload(w, item, r)
			return
		case both:
			http.Error(w, "Found view and download parameter only one can be supplied", http.StatusBadRequest)
			return
		case none:
			handleDirectory(w, item, r)
			return
		}
	}
	return
}
