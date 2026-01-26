// Created by claude as a Database test
package main

import (
	"context"
	"ferret/database"
	"fmt"
	"log"
)

func test() {
	// Connect to database
	url := "postgres://alize@localhost:5432/searchengine"
	conn := database.Connect(url)
	defer conn.Close(context.Background())

	fmt.Println("✅ Connected to database!")

	// Test 1: Insert a test file
	fmt.Println("\n📝 Inserting test data...")
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO files (filepath, filename, file_size, extension) VALUES ($1, $2, $3, $4)",
		"/test/sample.txt",
		"sample.txt",
		1024,
		".txt",
	)
	if err != nil {
		log.Fatal("Insert failed:", err)
	}
	fmt.Println("✅ Test data inserted!")

	// Test 2: Query all files
	fmt.Println("\n📖 Reading from database...")
	rows, err := conn.Query(
		context.Background(),
		"SELECT id, filepath, filename, file_size, extension FROM files",
	)
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

	// Test 3: Loop through and print results
	fmt.Println("\n📋 Files in database:")
	for rows.Next() {
		var id int
		var filepath, filename, extension string
		var fileSize int64

		err := rows.Scan(&id, &filepath, &filename, &fileSize, &extension)
		if err != nil {
			log.Fatal("Scan failed:", err)
		}

		fmt.Printf("ID: %d | %s | Size: %d bytes | Type: %s\n",
			id, filepath, fileSize, extension)
	}

	fmt.Println("\n🎉 Database test complete!")
}
