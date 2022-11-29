package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Bl4ckB3ard/golang-server-tool/config"
	"github.com/Bl4ckB3ard/golang-server-tool/dirparser"
	"github.com/Bl4ckB3ard/golang-server-tool/page"
	"github.com/Bl4ckB3ard/golang-server-tool/static"
	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if err := utils.CheckMethod(w, r); err != nil {
		return
	}

	content, _ := os.Open(config.ARGS.FilePath)

	utils.DefaultHeaders(w)

	http.ServeContent(w, r, filepath.Base(config.ARGS.FilePath), time.Now(), content)
}

func handleDirectory(w http.ResponseWriter, i dirparser.Item, r *http.Request) {
	if err := utils.CheckMethod(w, r); err != nil {
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

	utils.DefaultHeaders(w)
	tmpl.Execute(w, pageData)
	return
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if err := utils.CheckMethod(w, r); err != nil {
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

	utils.DefaultHeaders(w)

	tmpl.Execute(w, pageData)

	return
}

func handleView(w http.ResponseWriter, i dirparser.Item, r *http.Request) {
	if err := utils.CheckMethod(w, r); err != nil {
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
	if err := utils.CheckMethod(w, r); err != nil {
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
	if err := utils.CheckMethod(w, r); err != nil {
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
