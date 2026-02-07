package embedding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Model contains the configuration for an embedding service.
type Model struct {
	Name string
	Url  string
}

// Request represents the JSON structure for embedding API requests.
type Request struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// EmbeddingResponse represents the JSON structure of embedding API responses.
type EmbeddingResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
}

// Get sends text to the embedding service and returns the embedding vector.
// It returns an error if the HTTP request fails or no embeddings are returned.
func Get(embeddingModel Model, request Request) ([]float64, error) {

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
