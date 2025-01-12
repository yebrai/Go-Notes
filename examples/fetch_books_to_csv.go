package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type BookStruct struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type BooksRepo struct {
	Docs   []BookStruct `json:"docs"`
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
	Page   int          `json:"page"`
	Pages  int          `json:"pages"`
}

func main() {
	// Parse the filename flag
	filename := flag.String("filename", "books", "File name to save the books")
	flag.Parse()

	// Fetch books from the API
	resp, err := http.Get("https://the-one-api.dev/v2/book")
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Unmarshal the JSON data
	var books BooksRepo
	err = json.Unmarshal(body, &books)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Create the CSV file
	f, err := os.Create(*filename + ".csv")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer f.Close()

	// Write data to the CSV file
	w := csv.NewWriter(f)
	defer w.Flush()

	for _, book := range books.Docs {
		if err := w.Write([]string{book.Id, book.Name}); err != nil {
			log.Fatalf("Error writing to CSV: %v", err)
		}
	}

	// Check for errors in the CSV writer
	if err := w.Error(); err != nil {
		log.Fatalf("Error with CSV writer: %v", err)
	}

	fmt.Printf("Books saved to %s.csv\n", *filename)
}
