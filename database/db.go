package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pgvector/pgvector-go"
)

// Connect establishes a connection to PostgreSQL and registers the pgvector type.
// It exits the program if the connection fails.
func Connect(url string) *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	conn.TypeMap().RegisterType(&pgtype.Type{
		Name:  "vector",
		OID:   16385, // Vector type OID
		Codec: pgvector.VectorCodec{},
	})

	return conn
}

// InsertFileEmbedding inserts file metadata and its embedding vector into the database.
// It prints debug information and returns an error if the insertion fails.
func InsertFileEmbedding(conn *pgx.Conn, fileName, filePath, fileExtension, fileContents string, fileSize int, fileModifiedDate string, vector []float32) error {
	pgvector := pgvector.NewVector(vector)

	fmt.Println("===========================")
	fmt.Printf("FileName: %s\n", fileName)
	fmt.Printf("FilePath: %s\n", filePath)
	fmt.Printf("FileExtension: %s\n", fileExtension)
	fmt.Printf("FileSize: %d\n", fileSize)
	fmt.Printf("FileModifiedDate: %s\n", fileModifiedDate)
	fmt.Printf("Vector (first 60): %v...\n", vector[0:60])
	fmt.Println("===========================")

	_, err := conn.Exec(context.Background(),
		"INSERT INTO public.files (filename, filepath, content, file_size, modified_at, extension, embedding) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		fileName, filePath, fileContents, fileSize, fileModifiedDate, fileExtension, pgvector,
	)
	if err != nil {
		return fmt.Errorf("failed to insert embedding: %w", err)
	}
	return nil
}
