package database

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func Connect(url string) *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func vectorToString(vector []float64) string {
	// Convert []float64 to "[1.2, 3.4, 5.6, ...]" format
	parts := make([]string, len(vector))
	for i, v := range vector {
		parts[i] = fmt.Sprintf("%f", v)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func InsertFileEmbedding(conn *pgx.Conn, fileName string, filePath string, vector []float64) {
	vectorStr := vectorToString(vector)
	_, err := conn.Exec(context.Background(),
		"INSERT INTO public.files (filename, filepath, content, file_size, modified_at, extension, embedding) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		fileName, filePath, "", 0, "", "", vectorStr,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert embedding: %v\n", err)
	}
}
