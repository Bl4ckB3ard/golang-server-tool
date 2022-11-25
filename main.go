package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bl4ckB3ard/golang-server-tool/argparser"
	"github.com/Bl4ckB3ard/golang-server-tool/dirparser"
	"github.com/Bl4ckB3ard/golang-server-tool/handlers"
	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

func main() {

	if argparser.ARGS.IsFile {
		fmt.Printf("Serving file %v\nhttp://localhost:%s\n", argparser.ARGS.FilePath, argparser.ARGS.Port)

		http.HandleFunc("/", handlers.FileHandler)

		err := http.ListenAndServe(fmt.Sprintf(":%s", argparser.ARGS.Port), nil)
		utils.CheckServerErr(err)
	}
	// http.HandleFunc("/", handlers.MainHandler)
	// err := http.ListenAndServe(fmt.Sprintf(":%s", argparser.ARGS.Port), nil)
	// utils.CheckServerErr(err)
	names, err := dirparser.GetDirNames(argparser.ARGS.DirectoryPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(names)
}
