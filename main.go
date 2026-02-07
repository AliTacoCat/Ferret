package main

import (
	"context"
	"ferret/config"
	"ferret/database"
	"ferret/embedding"
	"ferret/walkdir"
	"fmt"
	"os"
)

func main() {
	// Load configuration
	cfg, err := config.Load("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Connect to database
	conn := database.Connect(cfg.Database.URL)
	defer conn.Close(context.Background())

	// Setup embedding model
	model := embedding.Model{
		Url:  cfg.Embedding.URL,
		Name: cfg.Embedding.ModelName,
	}

	// Walk files and get vector
	files := walkdir.GetAllFiles(cfg.FileWalk.RootPath, cfg.FileWalk)
	for _, file := range files {
		request := embedding.Request{
			Input: file.Contents,
			Model: model.Name,
		}

		embeddingVector, err := embedding.Get(model, request)
		if err != nil {
			panic(err)
		}

		vectorFloat32 := make([]float32, len(embeddingVector))
		for i, v := range embeddingVector {
			vectorFloat32[i] = float32(v) // Cast each element; watch for precision loss
		}

		// Store in database
		err = database.InsertFileEmbedding(conn, file.FileName, file.FullPath, file.Ext, file.Contents, int(file.Size), file.Modified, vectorFloat32)
		if err != nil {
			panic(err)
		}
	}
}
