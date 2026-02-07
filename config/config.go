package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds all application configuration parameters.
type Config struct {
	Database  DatabaseConfig  `json:"database"`
	Embedding EmbeddingConfig `json:"embedding"`
	FileWalk  FileWalkConfig  `json:"file_walk"`
}

// DatabaseConfig contains database connection settings.
type DatabaseConfig struct {
	URL string `json:"url"`
}

// EmbeddingConfig contains embedding service settings.
type EmbeddingConfig struct {
	URL       string `json:"url"`
	ModelName string `json:"model_name"`
}

// FileWalkConfig contains file walking and filtering settings.
type FileWalkConfig struct {
	RootPath         string   `json:"root_path"`
	SkipFolders      []string `json:"skip_folders"`
	SkipFiles        []string `json:"skip_files"`
	SkipExtensions   []string `json:"skip_extensions"`
}

// Load reads and parses the configuration file.
func Load(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}

// Save writes the configuration to a file.
func Save(cfg *Config, filepath string) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// Default returns a configuration with default values.
func Default() *Config {
	return &Config{
		Database: DatabaseConfig{
			URL: "postgres://alize@localhost:5432/searchengine",
		},
		Embedding: EmbeddingConfig{
			URL:       "http://localhost:1234/v1/embeddings",
			ModelName: "nomic-embed-text",
		},
		FileWalk: FileWalkConfig{
			RootPath: "/Users/alize/downloads",
			SkipFolders: []string{
				".", "node_modules", "vendor",
				"applications", "Library", ".git",
			},
			SkipFiles: []string{
				".", "thumbs.db",
			},
			SkipExtensions: []string{
				".jpg", ".mp4", ".png", ".dmg",
			},
		},
	}
}
