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
	Data []struct {
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
}

func Get(embeddingModel Model, request Request) ([]float64, error) {
	// LM Studio default endpoint
	// url := "http://localhost:1234/v1/embeddings"
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

	fmt.Println("Raw response from LM Studio:")
	fmt.Println(string(body))
	fmt.Println("---")

	// Deserialize JSON response into EmbeddingResponse struct
	var embeddingResp EmbeddingResponse
	if err := json.Unmarshal(body, &embeddingResp); err != nil {
		return nil, err
	}

	if len(embeddingResp.Data) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}
	return embeddingResp.Data[0].Embedding, nil
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
