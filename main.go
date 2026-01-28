package main

import (
	"context"
	"ferret/database"
	"ferret/embedding"
	"ferret/walkdir"
)

func main() {

	//connects to database
	dbURL := "postgres://alize@localhost:5432/searchengine"
	conn := database.Connect(dbURL)
	defer conn.Close(context.Background())

	//walk and embed files
	embeddingModel := embedding.Model{
		Url:  "http://localhost:1234/v1/embeddings",
		Name: "nomic-embed-text",
	}

	files := walkdir.GetAllFiles("/Users/alize/downloads")
	for _, file := range files {
		request := embedding.Request{
			Input: file.Contents,
			Model: embeddingModel.Name,
		}

		embeddingVector, err := embedding.Get(embeddingModel, request)
		if err != nil {
			panic(err)
		}

		//store in database
		database.InsertFileEmbedding(conn, file.FileName, file.FullPath, embeddingVector)
	}

}
