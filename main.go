package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bl4ckB3ard/golang-server-tool/config"
	"github.com/Bl4ckB3ard/golang-server-tool/handlers"
)


func main() {
    fmt.Println(config.ARGS)

	if config.ARGS.IsFile {
		fmt.Printf("Serving file %v\nhttp://localhost:%s\n", config.ARGS.FilePath, config.ARGS.Port)

		http.HandleFunc("/", handlers.FileHandler)

		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ARGS.Port), nil)
        if err != nil {
            fmt.Fprintf(os.Stderr, "ERROR Starting server\n")
            return
        }
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandler)
    mux.HandleFunc("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/", handlers.StaticHandler)

	fmt.Printf("Serving %v\nhttp://localhost:%s\n", config.ARGS.DirectoryPath, config.ARGS.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ARGS.Port), mux)

    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR Starting server\n")
        return }
}
