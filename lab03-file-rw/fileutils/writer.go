package fileutils

import "os"

func WriteTextFile(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
