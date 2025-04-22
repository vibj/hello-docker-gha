package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define a struct to match the JSON response structure
type Post struct {
	Title string `json:"title"`
}

func main() {
	fmt.Println("Starting to call the API...")

	// Define the URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v", err)
	}

	// Parse the JSON response
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v", err)
	}

	// Print the title to the logs
	fmt.Printf("Title: %s", post.Title)
}
