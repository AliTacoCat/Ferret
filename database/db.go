package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pgvector/pgvector-go"
)

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

func InsertFileEmbedding(conn *pgx.Conn, fileName, filePath, fileExtension, fileContents string, fileSize int, fileModifiedDate string, vector []float32) {
	pgvector := pgvector.NewVector(vector)
	println("===========================")
	println("FileName:", fileName)
	println("FilePath:", filePath)
	println("FileExtension:", fileExtension)
	println("FileSize:", fileSize)
	println("FileModifiedDate:", fileModifiedDate)
	println("VectorStr:", vector[0:60], "...")
	println("===========================")

	_, err := conn.Exec(context.Background(),
		"INSERT INTO public.files (filename, filepath, content, file_size, modified_at, extension, embedding) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		fileName, filePath, fileContents, fileSize, fileModifiedDate, fileExtension, pgvector,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert embedding: %v\n", err)
	}
}
