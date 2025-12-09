package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EmbeddingModel struct {
	Name string
	Url  string
}

type EmbeddingRequest struct {
<<<<<<< HEAD
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

type EmbeddingResponse struct {
	Embedding []float64 `json:"embedding"`
=======
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
>>>>>>> 6f34e653d8d8ba5c07e44a4105e180fe46bfe457
}

func GetEmbedding(text string) ([]float64, error) {
	// LM Studio default endpoint
	url := "http://localhost:11434/api/embeddings"
	fmt.Println("Using URL:", url)

	reqBody := EmbeddingRequest{
<<<<<<< HEAD
		Prompt: text,
		Model:  "nomic-embed-text", // Replace with your model name
=======
		Input: text,
		Model: "nomic-embed-text", // Replace with your model name
>>>>>>> 6f34e653d8d8ba5c07e44a4105e180fe46bfe457
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

<<<<<<< HEAD
	fmt.Println("Raw response from Ollama:")
	fmt.Println(string(body))
	fmt.Println("---")

=======
>>>>>>> 6f34e653d8d8ba5c07e44a4105e180fe46bfe457
	// Deserialize JSON response into EmbeddingResponse struct
	var embeddingResp EmbeddingResponse
	if err := json.Unmarshal(body, &embeddingResp); err != nil {
		return nil, err
	}

	// Assert that we have at least one embedding returned
<<<<<<< HEAD
	if len(embeddingResp.Embedding) == 0 {
=======
	if len(embeddingResp.Data) == 0 {
>>>>>>> 6f34e653d8d8ba5c07e44a4105e180fe46bfe457
		return nil, fmt.Errorf("no embeddings returned")
	}

	// Return the first embedding vector
<<<<<<< HEAD
	return embeddingResp.Embedding, nil
=======
	return embeddingResp.Data[0].Embedding, nil
>>>>>>> 6f34e653d8d8ba5c07e44a4105e180fe46bfe457
}

func main() {
	text := "Hello, this is a test sentence for embedding generation."

	embedding, err := GetEmbedding(text)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("First 5 values: %v\n", embedding[:5])
	fmt.Printf("The Length of embedding is %d\n", len(embedding))
}
