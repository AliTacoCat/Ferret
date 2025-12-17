package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Files struct {
	FullPath string
	FileName string
	Size     int64
	Modified string
	ext      string
	isFolder bool
	Contents string
}

func GetAllFiles(path string) []Files {
	var listOfFiles []Files
	FoldSkip := []string{".", "node_modules", "vendor", "applications", "Library", ".git"}
	FileSkip := []string{".", "thumbs.db", ".jpg", ".mp4", ".png", ".dmg"}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		file := Files{
			FullPath: path,
			FileName: info.Name(),
			Size:     info.Size(),
			Modified: info.ModTime().String(),
			ext:      filepath.Ext(info.Name()),
			isFolder: info.IsDir(),
			Contents: "",
		}

		if contains(FoldSkip, info.Name()) || contains(FileSkip, info.Name()) {
			return nil
		}
		if contains(FileSkip, file.ext) {
			return nil
		}

		listOfFiles = append(listOfFiles, file)
		return nil
	})

	return listOfFiles
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// func marshalFilesToJSON(files []Files, prompt string) marshalFiles {
// 	return marshalFiles{
// 		model:  files,
// 		prompt: prompt,
// 	}
// }

func run() {
	files := GetAllFiles("/Users/alize/downloads")

	for _, file := range files {
		fmt.Println(file.FileName)
		fmt.Println("Path:", file.FullPath)
		fmt.Println("Size:", file.Size)
		fmt.Println("Modified:", file.Modified)
		fmt.Println("Extension:", file.ext)
		fmt.Println("Is Folder:", file.isFolder)
		if file.ext == ".txt" || file.ext == ".rtf" {
			fmt.Println("Contents:")
			content, err := os.ReadFile(file.FullPath)
			if err != nil {
				fmt.Println("Error reading file:", err)
			} else {
				fmt.Println(string(content))
			}
		}
		fmt.Println("-----")
	}
	fmt.Println("Total files:", len(files))
}
