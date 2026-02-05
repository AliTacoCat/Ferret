package main

import (
	"context"
	"ferret/database"
	"ferret/embedding"
	"ferret/walkdir"
)

func main() {

	// Connects to database
	dbURL := "postgres://alize@localhost:5432/searchengine"
	conn := database.Connect(dbURL)
	defer conn.Close(context.Background())

	// Walk and embed files
	model := embedding.Model{
		Url:  "http://localhost:1234/v1/embeddings",
		Name: "nomic-embed-text",
	}

	// Walk files and get vector
	files := walkdir.GetAllFiles("/Users/alize/downloads")
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
