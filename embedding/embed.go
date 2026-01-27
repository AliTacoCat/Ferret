package embedding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Model struct {
	Name string
	Url  string
}

type Request struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type EmbeddingResponse struct {
	Embedding []float64 `json:"embedding"`
}

func Get(embeddingModel Model, request Request) ([]float64, error) {
	// LM Studio default endpoint
	// url := "http://localhost:11434/api/embeddings"
	fmt.Println("Using URL:", embeddingModel.Url)

	// Serialize request body to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Send JSON HTTP POST request and wait for response
	resp, err := http.Post(embeddingModel.Url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body into byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Raw response from Ollama:")
	fmt.Println(string(body))
	fmt.Println("---")

	// Deserialize JSON response into EmbeddingResponse struct
	var embeddingResp EmbeddingResponse
	if err := json.Unmarshal(body, &embeddingResp); err != nil {
		return nil, err
	}

	// Assert that we have at least one embedding returned
	if len(embeddingResp.Embedding) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}

	// Return the first embedding vector
	return embeddingResp.Embedding, nil
}

// func TestEmbed() {
// 	text := "Hello, this is a test sentence for embedding generation."

// 	embedding, err := GetEmbedding(text)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("First 5 values: %v\n", embedding[:5])
// 	fmt.Printf("The Length of embedding is %d\n", len(embedding))
// }
