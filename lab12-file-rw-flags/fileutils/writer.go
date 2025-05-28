package fileutils

import "os"

func WriteFileContent(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
