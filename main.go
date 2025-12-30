package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
)

// URLShortener holds our database (a map) and a Mutex.
// Maps in Go are not thread-safe by default, so we use sync.Mutex
// to make sure two people don't try to write to the map at the same time.
type URLShortener struct {
	urls map[string]string
	mu   sync.Mutex
}

func main() {
	shortener := &URLShortener{
		urls: make(map[string]string),
	}

	// Route 1: Shorten a URL
	// Usage: http://localhost:8080/shorten?url=https://google.com
	http.HandleFunc("/shorten", shortener.handleShorten)

	// Route 2: Redirect using the code
	// Usage: http://localhost:8080/abc123
	http.HandleFunc("/", shortener.handleRedirect)

	fmt.Println("URL Shortener Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// generateCode creates a random 6-character string.
func generateCode() string {
	b := make([]byte, 3) // 3 bytes = 6 hex characters
	if _, err := rand.Read(b); err != nil {
		return "default"
	}
	return hex.EncodeToString(b)
}

func (s *URLShortener) handleShorten(w http.ResponseWriter, r *http.Request) {
	// Get the "url" parameter from the query string
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	// Generate code and save to map
	code := generateCode()
	
	s.mu.Lock()
	s.urls[code] = longURL
	s.mu.Unlock()

	// Respond to the user
	shortURL := fmt.Sprintf("http://localhost:8080/%s", code)
	fmt.Fprintf(w, "Shortened URL: %s\n", shortURL)
}

func (s *URLShortener) handleRedirect(w http.ResponseWriter, r *http.Request) {
	// The path will be "/code", so we skip the first character "/"
	code := r.URL.Path[1:]

	s.mu.Lock()
	originalURL, exists := s.urls[code]
	s.mu.Unlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Perform the redirect
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}	