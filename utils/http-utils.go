package utils

import (
	"errors"
	"fmt"
	"net/http"
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
