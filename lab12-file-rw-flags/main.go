package main

import (
	"fmt"
	"log"
	"os"
	"pattern-decorator/fileutils"
	"reflect"
)

func main() {
	// log.SetFlags(0)
	log.SetPrefix("[MegaRW] ")

	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	// TEST READING
	filePath := (rootPath + "/data/reading.txt")

	log.Printf("root path: %s \n", rootPath)
	log.Printf("file path: %s \n", filePath)
	log.Printf("file path type: %s \n", reflect.TypeOf(filePath))

	c, err := fileutils.ReadFileContent(filePath, fileutils.WithLineNums|fileutils.WithBar)

	if err != nil {
		log.Fatalf("Error! Can't Read!\n %v", err)
	}

	fmt.Printf("%s\n", c)

	// TEST WRITING
	filePath = (rootPath + "/data/writing.txt")

	newContent := fmt.Sprintf("This is cool!\nAnother line? WOW\nReally! Now Jump!\n\nDAFU@#()!!!\n")

	err = fileutils.WriteFileContent(filePath, newContent)

	if err != nil {
		log.Fatalf("Error! Can't Write!\n %v", err)
	}

	log.Printf("Wrote!\n")

}
