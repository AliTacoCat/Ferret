package walkdir

import (
	"ferret/config"
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
// for all non-skipped files. It uses the provided config for filtering rules.
func GetAllFiles(path string, cfg config.FileWalkConfig) []Files {
	var listOfFiles []Files

	// Use maps for O(1) lookup instead of O(n) slice search
	foldSkip := make(map[string]bool)
	for _, folder := range cfg.SkipFolders {
		foldSkip[folder] = true
	}

	fileSkip := make(map[string]bool)
	for _, file := range cfg.SkipFiles {
		fileSkip[file] = true
	}

	extSkip := make(map[string]bool)
	for _, ext := range cfg.SkipExtensions {
		extSkip[ext] = true
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
// It loads configuration from config.json.
func WalkEmbed() {
	cfg, err := config.Load("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	model := embedding.Model{
		Url:  cfg.Embedding.URL,
		Name: cfg.Embedding.ModelName,
	}

	files := GetAllFiles(cfg.FileWalk.RootPath, cfg.FileWalk)

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
