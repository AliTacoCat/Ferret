package main

import (
	"context"
	"ferret/database"
	"ferret/embedding"
	"ferret/walkdir"
	"os"
)

func main() {
	// Step 1: Read files from directory
	files := walkdir.GetAllFiles("/Users/alize/downloads")

	// Step 2: Connect to database
	url := "postgres://alize@localhost:5432/searchengine"
	conn := database.Connect(url)
	defer conn.Close(context.Background())

	// Step 3: Insert file metadata into database
	for _, file := range files {
		_, err := conn.Exec(
			context.Background(),
			"INSERT INTO files (filename, file_size, extension) VALUES ($1, $2, $3)",
			file.FileName,
			file.Size,
			file.ext,
		)
		if err != nil {
			panic(err)
		}
	}

	// Step 4: Generate embeddings for text files and store them
	embeddingModel := embedding.Model{
		Name: "nomic-embed-text",
		Url:  "http://localhost:11434/api/embeddings",
	}

	for _, file := range files {
		if file.ext == ".txt" && file.Contents != "" {
			request := embedding.Request{
				Input: file.Contents,
				Model: embeddingModel.Name,
			}

			embeddingVector, err := embedding.Get(embeddingModel, request)
			if err != nil {
				panic(err)
			}

			// Store embedding in database (assuming a table 'embeddings' exists)
			_, err = conn.Exec(
				context.Background(),
				"INSERT INTO embeddings (filename, embedding) VALUES ($1, $2)",
				file.FileName,
				embeddingVector,
			)
			if err != nil {
				panic(err)
			}
		}
	}

	os.Exit(0)

}
