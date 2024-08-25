package main

import (
	"file-rw/fileutils"
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/sample.txt"

	c, err := fileutils.ReadTextFile(filePath)

	if err == nil {
		fmt.Println(c)

		newContent := fmt.Sprintf("Original: %v\n\nDouble Original:\n%v%v", c, c, c)
		fileutils.WriteTextFile(filePath, newContent)

	} else {
		fmt.Printf("PANIC! %v", err)
	}

}
