package fileutils

import "os"

func ReadTextFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err == nil {
		return string(content), nil
	} else {
		return "", err
	}
}
