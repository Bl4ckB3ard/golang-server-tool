package utils

import (
	"fmt"
	"os"
	"strconv"
)

func IsValidPort(port string) bool {
	num, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return false
	}
	return num >= 1 && num < 65536
}

func CheckServerErr(err error) {
	if err != nil {
		fmt.Println("ERROR starting server")
		os.Exit(1)
	}
}
