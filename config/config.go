package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Bl4ckB3ard/golang-server-tool/dirparser"
	"github.com/akamensky/argparse"
)

type Args struct {
	DirectoryPath string
	Port          string
	FilePath      string
	IsFile        bool
	Theme         string
}

func (a* Args)Validate() {
    // this function will exit if args are incorrect
    if a.IsFile {
        if !isValidFile(a.FilePath) {
            fmt.Printf("ERROR %v is not a valid file\n", a.FilePath)
            os.Exit(1)
            return
        }
    }

    if !a.IsFile {
        if !isValidDirectory(a.DirectoryPath) {
            fmt.Printf("ERROR %v is not a valid directory\n", a.FilePath)
            os.Exit(1)
            return
        }
    }

    if !isValidPort(a.Port) {
            fmt.Printf("ERROR %v is not a valid port\n", a.Port)
            os.Exit(1)
            return
    }

    if !isValidTheme(a.Theme) {
        fmt.Printf("ERROR %v is not a valid theme\n", a.Theme)
        os.Exit(1)
        return
    }

    return
}


func isValidTheme(theme string) bool {
    validThemes := [2]string{"light", "dark"}

    for _, i := range(validThemes) {
        if strings.ToLower(theme) == i {
            return true
        }
    }

    return false
}


func isValidDirectory(dir string) bool {
	fullPath, err := filepath.Abs(dir)
	if err != nil {
		return false
	}
	fileHandle, err := os.Open(fullPath)
	defer fileHandle.Close()
	if err != nil {
		return false
	}
	info, err := fileHandle.Stat()
	if err != nil {
		return false
	}
	return info.IsDir()
}


func isValidFile(file string) bool {
	fullPath, err := filepath.Abs(file)
	if err != nil {
		return false
	}
	fileHandle, err := os.Open(fullPath)
	defer fileHandle.Close()
	if err != nil {
		return false
	}
	info, err := fileHandle.Stat()
	if err != nil {
		return false
	}
	return !info.IsDir()
}


func isValidPort(port string) bool {
	num, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return false
	}
	return num >= 1 && num < 65536
}


var (
	ARGS           Args
	RootFS         dirparser.RootFS
)


func init() {
    var isFile bool
    parser := argparse.NewParser("Golang-Server-Tool", "cmd-line http server for serving local files.")
    dirPath := parser.String("d", "dir", &argparse.Options{Required: false, Help: "Path to directory to serve files from."})
    filePath := parser.String("f", "file", &argparse.Options{Required: false, Help: "Path to a file to serve only a single file."})
    port := parser.String("p", "port", &argparse.Options{Required: false, Help: "TCP port for the server to run on", Default: "8080"})
    theme := parser.String("t", "theme", &argparse.Options{Required: false, Help: "Dark or Light theme for the HTML", Default: "light"})
    err := parser.Parse(os.Args)
    

    if err != nil {
        fmt.Println(parser.Usage(err))
        return
    }

    // get cwd if no args supplied
    if *dirPath == "" && *filePath == "" {
        cwd, err := os.Getwd()

        if err != nil {
            fmt.Fprintf(os.Stderr, "ERROR getting current working directory\n")
            return
        }

        *dirPath = cwd
    }

    if *dirPath != "" && *filePath != "" { // if both -d and -f supplied
        fmt.Printf("BOTH -d and -f supplied you can only use one\n")
        return
    }

    isFile = *filePath != ""

    ARGS = Args{
        DirectoryPath: *dirPath,
        Port: *port,
        FilePath: *filePath,
        IsFile: isFile,
        Theme: *theme,
    }

    ARGS.Validate()

    if ! ARGS.IsFile {
        RootFS, err = dirparser.GetRootFS(ARGS.DirectoryPath)

        if err != nil {
            fmt.Println(err)
            os.Exit(1)
            return
        }

    }


    return 
}
