package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadTxt(file os.DirEntry) {
	if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
		fmt.Println("\n=== Reading: ", file.Name(), " ===")

		fullpath := filepath.Join("/Users/alize/downloads", file.Name())
		content, err := os.ReadFile(fullpath)

		if err != nil {
			fmt.Println("Error reading file:", file.Name(), ":", err)
		} else {
			fmt.Printf("Content:\n%s\n", string(content))
		}
		return
	}
}

func ReadDir() (int, int, map[string]int) {
	foldercount := 0
	filecount := 0
	counts := make(map[string]int)
	files, err := os.ReadDir("/Users/alize/downloads")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return foldercount, filecount, counts
	}
	for _, file := range files {
		info, _ := file.Info()
		ReadTxt(file)

		ext := filepath.Ext(file.Name())
		counts[ext]++
		if file.IsDir() {
			fmt.Println("\n", file.Name(), "(directory)")
			foldercount++
		} else {
			fmt.Println("\n", file.Name(), "(file)")
			fmt.Printf("Size: %d bytes \n", info.Size())
			fmt.Printf("Modified: %s \n", info.ModTime())
			filecount++
		}
	}
	return foldercount, filecount, counts
}

func PrintDir() {
	foldercount, filecount, counts := ReadDir()
	fmt.Printf("\n Total files: %d\n", filecount)
	fmt.Println("\n File types found:")

	for ext, fileCount := range counts {
		fmt.Printf("Extension: '%s' - Count: %d\n", ext, fileCount)
	}
	fmt.Printf("Total directories: %d\n", foldercount)
}
