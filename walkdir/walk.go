package walkdir

import (
	"ferret/embedding"
	"fmt"
	"os"
	"path/filepath"
)

// Files represents metadata about a file in the filesystem.
type Files struct {
	FullPath string
	FileName string
	Size     int64
	Modified string
	Ext      string
	isFolder bool
	Contents string
}

// GetAllFiles walks the directory tree starting at path and returns metadata
// for all non-skipped files. It filters out common directories (node_modules,
// vendor, .git) and file types (.jpg, .mp4, .png, .dmg).
func GetAllFiles(path string) []Files {
	var listOfFiles []Files

	// Use maps for O(1) lookup instead of O(n) slice search
	foldSkip := map[string]bool{
		".": true, "node_modules": true, "vendor": true,
		"applications": true, "Library": true, ".git": true,
	}
	fileSkip := map[string]bool{
		".": true, "thumbs.db": true,
	}
	extSkip := map[string]bool{
		".jpg": true, ".mp4": true, ".png": true, ".dmg": true,
	}

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
			Ext:      filepath.Ext(info.Name()),
			isFolder: info.IsDir(),
			Contents: "",
		}

		if foldSkip[info.Name()] || fileSkip[info.Name()] || extSkip[file.Ext] {
			return nil
		}

		listOfFiles = append(listOfFiles, file)
		return nil
	})

	return listOfFiles
}

// WalkEmbed is a debug function that walks files and attempts to generate
// embeddings for .txt and .rtf files, printing results to stdout.
func WalkEmbed() {
	model := embedding.Model{
		Url:  "http://localhost:1234/v1/embeddings",
		Name: "nomic-embed-text",
	}

	files := GetAllFiles("/Users/alize/downloads")

	for _, file := range files {
		request := embedding.Request{
			Input: file.Contents,
			Model: model.Name,
		}
		fmt.Println(file.FileName)
		fmt.Println("Path:", file.FullPath)
		fmt.Println("Size:", file.Size)
		fmt.Println("Modified:", file.Modified)
		fmt.Println("Extension:", file.Ext)
		fmt.Println("Is Folder:", file.isFolder)
		if file.Ext == ".txt" || file.Ext == ".rtf" {
			fmt.Println("Contents:")
			content, err := os.ReadFile(file.FullPath)
			if err != nil {
				fmt.Println("Error reading file:", err)
			} else {
				fmt.Println(string(content))
			}
			vector, err := embedding.Get(model, request)
			if err != nil {
				fmt.Println("Error getting embedding:", err)
			} else {
				fmt.Println("Embedding vector:", vector)
			}
		}
		fmt.Println("-----")
	}
	fmt.Println("Total files:", len(files))
}

// store name, path, any identifible information about the file after getting embeddings
