package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileStruct struct {
	Name     string
	Size     int64
	Modified time.Time
	ext      string
	isFolder bool
	Content  string
}

func ReadTxt(file os.DirEntry) (content string) {
	if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {

		fullpath := filepath.Join("/Users/alize/downloads", file.Name())
		content, err := os.ReadFile(fullpath)

		if err != nil {
			fmt.Println("Error reading file:", file.Name(), ":", err)
		} else {
			return string(content)
		}
	}
	return
}

func ReadDir() (map[string]int, []FileStruct) {
	counts := make(map[string]int)
	files, err := os.ReadDir("/Users/alize/downloads")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return counts, nil
	}

	var StructList []FileStruct

	for _, file := range files {
		info, _ := file.Info()

		ext := filepath.Ext(file.Name())
		counts[ext]++
		returnedFileInfo := FileStruct{
			Name:     info.Name(),
			Size:     info.Size(),
			Modified: info.ModTime(),
			ext:      ext,
			isFolder: file.IsDir(),
			Content:  ReadTxt(file),
		}

		for _, fileInfo := range StructList {
			if fileInfo.Name == returnedFileInfo.Name {
				continue
			}
		}
		StructList = append(StructList, returnedFileInfo)
	}
	return counts, StructList
}

func PrintDir() {

	counts, StructList := ReadDir()
	////
	fmt.Println("\n Detailed File Information:")
	for _, fileInfo := range StructList {
		if fileInfo.ext == ".txt" && !fileInfo.isFolder {
			fmt.Printf("\n Text File: %s \n Size: %d bytes \n Modified: %s \n Extension: %s\n Content: \n%s\n",
				fileInfo.Name, fileInfo.Size, fileInfo.Modified.Format(time.RFC1123), fileInfo.ext, fileInfo.Content)
		} else if fileInfo.isFolder {
			fmt.Printf("\n Directory: %s \n | Modified: %s\n", fileInfo.Name+"\n", fileInfo.Modified.Format(time.RFC1123))
		} else {
			fmt.Printf("\n File: %s \n Size: %d bytes \n Modified: %s \n Extension: %s\n",
				fileInfo.Name, fileInfo.Size, fileInfo.Modified.Format(time.RFC1123), fileInfo.ext)
		}
	}

	fmt.Println("\n File types found:")

	for ext, fileCount := range counts {
		fmt.Printf("Extension: '%s' - Count: %d\n", ext, fileCount)
	}
}
