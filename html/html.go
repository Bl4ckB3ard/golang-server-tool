package html

import (
	"fmt"
	"os"
)

func createTempFile(filePath string, content []byte) {
	fileHandle, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("ERROR: couldn't create file %v\n", filePath)
		os.Exit(2)
	}
	fileHandle.Write(content)
}

// TODO:
func createHtml(itemNames []string, theme string) error {
	var html string

	if theme == "" {
		theme = "normal"
	}

	itemsInDir := len(itemNames)

}
