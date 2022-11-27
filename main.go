package main

import (
	"fmt"
	"net/http"

	"github.com/Bl4ckB3ard/golang-server-tool/config"
	"github.com/Bl4ckB3ard/golang-server-tool/handlers"
	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

func main() {
	if config.ARGS.IsFile {
		fmt.Printf("Serving file %v\nhttp://localhost:%s\n", config.ARGS.FilePath, config.ARGS.Port)

		http.HandleFunc("/", handlers.FileHandler)

		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ARGS.Port), nil)
		utils.CheckServerErr(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/favicon.ico", handlers.FaviconHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ARGS.Port), mux)
	utils.CheckServerErr(err)
}
