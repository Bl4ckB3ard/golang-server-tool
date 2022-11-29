package main

import (
	"fmt"
	"net/http"

	"github.com/Bl4ckB3ard/golang-server-tool/config"
	"github.com/Bl4ckB3ard/golang-server-tool/handlers"
	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

func SetupStaticRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/favicon.ico", handlers.FaviconHandler)
	// light
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/file_icon_light.png", handlers.FileLight)
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/folder_icon_light.png", handlers.FolderLight)
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/sort_arrow_icon_light.png", handlers.ArrowLight)
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/logo_light.png", handlers.LogoLight)
	// dark
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/file_icon_dark.png", handlers.FileDark)
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/folder_icon_dark.png", handlers.FolderDark)
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/sort_arrow_icon_dark.png", handlers.ArrowDark)
	mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/logo_dark.png", handlers.LogoDark)

}

func main() {
	if config.ARGS.IsFile {
		fmt.Printf("Serving file %v\nhttp://localhost:%s\n", config.ARGS.FilePath, config.ARGS.Port)

		http.HandleFunc("/", handlers.FileHandler)

		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ARGS.Port), nil)
		utils.CheckServerErr(err)
	}

	mux := http.NewServeMux()
	SetupStaticRoutes(mux)

	mux.HandleFunc("/", handlers.MainHandler)

	fmt.Printf("Serving %v\nhttp://localhost:%s\n", config.ARGS.DirectoryPath, config.ARGS.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ARGS.Port), mux)
	utils.CheckServerErr(err)
}
