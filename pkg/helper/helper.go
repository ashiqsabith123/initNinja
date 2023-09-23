package helper

import (
	"fmt"

	"os"
	"path/filepath"
	"time"

	"github.com/ashiqsabith123/initNinja/pkg/ui"
)

func CreateFolder(folder_name string, directories []string) {

	status := make(chan string)

	projectRoot := filepath.Join(".", folder_name)

	go func() {
		for _, dir := range directories {
			dirPath := filepath.Join(projectRoot, dir)
			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dirPath, err)
				return
			}
		}
		time.Sleep(3 * time.Second)

		status <- "C"
	}()

	ui.Spinner(35, 3, "Intializing folder", "magenta", status)

}

func CreateFiles(folder_name string, file_names []string) {

	for _, v := range file_names {

		filepath := folder_name + "/" + v

		file, err := os.Create(filepath)

		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
	}
}
