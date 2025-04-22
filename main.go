package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	// Write the title to the GITHUB_ENV file
	githubEnv := os.Getenv("GITHUB_ENV")
	if githubEnv != "" {
		file, err := os.OpenFile(githubEnv, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Failed to open GITHUB_ENV file: %v", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("TITLE=%s\n", post.Title))
		if err != nil {
			fmt.Printf("Failed to write to GITHUB_ENV file: %v", err)
		}
	} else {
		fmt.Printf("GITHUB_ENV is not set. Running outside of GitHub Actions?")
	}
}
