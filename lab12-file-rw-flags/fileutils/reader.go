package fileutils

import (
	"fmt"
	"os"
	"strings"
)

const (
	NoFlags      = 0
	WithLineNums = 1 << iota
	WithBar
)

func ReadFileContent(filePath string, flags int) (string, error) {

	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	if flags&WithLineNums == 0 {
		return string(content), nil
	}

	lines := strings.Split(string(content), "\n")
	numLines := []string{}

	for i, line := range lines {
		numLines = append(numLines, fmt.Sprintf("%4d: %s", i+1, line))
	}

	if flags&WithBar != 0 {
		numLines = append(numLines, "-----------")
	}

	numStr := strings.Join(numLines, "\n")

	return numStr, nil

}
