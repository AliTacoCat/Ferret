package main

// func ignore() {
// 	// Step 1: Read all files from the specified directory
// 	path := "/Users/alize/downloads"
// 	files := walkdir.GetAllFiles(path)

// 	// Step 2: Connect to the database
// 	dbURL := "postgres://alize@localhost:5432/searchengine"
// 	conn := database.Connect(dbURL)
// 	defer conn.Close(context.Background())

// 	// Step 3: Initialize embedding model
// 	embeddingModel := embedding.Model{
// 		Url:  "http://localhost:11434/api/embeddings",
// 		Name: "nomic-embed-text",
// 	}

// 	// Step 4: Process each file
// 	for _, file := range files {
// 		// Get embedding for the file content
// 		request := embedding.Request{
// 			Input: file.Contents,
// 			Model: embeddingModel.Name,
// 		}
// 		vector, err := embedding.Get(embeddingModel, request)
// 		if err != nil {
// 			continue
// 		}

// 		//	 Insert file metadata and embedding into the databas
// 		_, err = conn.Exec(
// 			context.Background(),
// 			"INSERT INTO files (filepath, filename, file_size, extension, embedding) VALUES ($1, $2, $3, $4, $5)",
// 			path,
// 			file.FileName,
// 			file.Size,
// 			file.Ext,
// 			vector,
// 		)
// 		if err != nil {
// 			continue
// 		}
// 	}

// 	os.Exit(0)
// }
