package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolder(folder_name string, directories []string) {

	projectRoot := filepath.Join(".", folder_name)

	for _, dir := range directories {
		dirPath := filepath.Join(projectRoot, dir)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dirPath, err)
			return
		}
		fmt.Printf("Created directory: %s\n", dirPath)
	}
}
