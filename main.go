package main

import (
	"context"
	"ferret/database"
	"os"
)

func main() {
	conn := database.Connect(os.Getenv("postgres://Alize:@localhost:5432/postgres"))
	defer conn.Close(context.Background())

	// var name string
	// var weight int64
	// err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(name, weight)

	embeddingModel := EmbeddingModel{
		Name: "nomic-embed-text",
		Url:  "http://localhost:11434/api/embeddings",
	}

	embeddingRequest := EmbeddingRequest{
		Input: "Hello, this is a test sentence for embedding generation.",
		Model: "nomic-embed-text",
	}

	_, err := GetEmbedding(embeddingModel, embeddingRequest)
	if err != nil {
		panic(err)
	}

}
