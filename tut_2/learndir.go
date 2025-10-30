package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 1. Establish a connection to the database
// 		1. Where do we store it?
// 2. Walk the directories
// 		1. Avoid System dirctories
// 3. Open and Index files we care about
// 		1. Store the INDEX of the files info in the DB
// 		2. Dont forget the meta data
// 4. Have a way to search the stored index
// 		1. Needs to return the path of the file
// 		2. Should use filters to select certain file types

type Files struct {
	FullPath string
	FileName string
}

func GetAllFiles(path string) []Files {
	var listOfFiles []Files

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == true {
			return nil
		}

		file := Files{
			FullPath: path,
			FileName: info.Name(),
		}

		listOfFiles = append(listOfFiles, file)
		return nil
	})

	return listOfFiles
}

func main() {
	files := GetAllFiles(".")

	for _, file := range files {
		fmt.Println(file)
	}
}
