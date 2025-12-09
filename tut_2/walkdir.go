package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type EmbeddingModel struct {
	Name string
	Url  string
}

type EmbeddingRequest struct {
	Input string `json:"prompt"`
	Model string `json:"model"`
}

type EmbeddingResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

func GetEmbedding(text string) ([]float64, error) {
	// LM Studio default endpoint
	url := "http://localhost:11434/api/embeddings"

	reqBody := EmbeddingRequest{
		Input: text,
		Model: "nomic-embed-text", // Replace with your model name
	}

	// Serialize request body to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Send JSON HTTP POST request and wait for response
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body into byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Deserialize JSON response into EmbeddingResponse struct
	var embedResp EmbeddingResponse
	err = json.Unmarshal(body, &embedResp)
	if err != nil {
		return nil, err
	}

	if len(embedResp.Data) > 0 {
		return embedResp.Data[0].Embedding, nil
	}

	return nil, fmt.Errorf("no embedding data found in response")
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
