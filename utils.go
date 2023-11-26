package tokens

import (
	"fmt"
	"os"
	"path/filepath"
)

func getJSONFiles() []string {
	var list []string
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return list
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return list
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".json" {
			list = append(list, filepath.Join(dir, file.Name()))
		}
	}

	return list
}