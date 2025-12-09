package tut_2

import (
	"fmt"
	"os"
	"path/filepath"
<<<<<<< HEAD
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
=======
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
>>>>>>> 6f34e653d8d8ba5c07e44a4105e180fe46bfe457
}
