package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Bl4ckB3ard/golang-server-tool/argparser"
	"github.com/Bl4ckB3ard/golang-server-tool/utils"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if err := utils.CheckMethod(w, r); err != nil {
		return
	}
	content, _ := os.Open(argparser.ARGS.FilePath)
	utils.DefaultHeaders(w)
	http.ServeContent(w, r, filepath.Base(argparser.ARGS.FilePath), time.Now(), content)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {

}
