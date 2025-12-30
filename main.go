package main

import (
	"database/sql"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the driver
)

// Link represents our data model
type Link struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	Clicks      int    `json:"clicks"`
}

var db *sql.DB

func main() {
	// 1. Initialize Database
	var err error
	db, err = sql.Open("sqlite3", "./links.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if it doesn't exist
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS links (id INTEGER PRIMARY KEY, original_url TEXT, short_code TEXT, clicks INTEGER)")
	statement.Exec()

	// 2. Routes
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", handleRedirect)

	fmt.Println("Pro URL Shortener started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateCode() string {
	b := make([]byte, 3)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request
	var input struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	code := generateCode()

	// Insert into Database
	_, err := db.Exec("INSERT INTO links (original_url, short_code, clicks) VALUES (?, ?, ?)", input.URL, code, 0)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", code),
	})
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	if code == "" {
		fmt.Fprint(w, "Welcome to the Pro Shortener!")
		return
	}

	var originalURL string
	var clicks int

	// Query Database & Update Clicks
	err := db.QueryRow("SELECT original_url, clicks FROM links WHERE short_code = ?", code).Scan(&originalURL, &clicks)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	// Increment click counter in background
	db.Exec("UPDATE links SET clicks = ? WHERE short_code = ?", clicks+1, code)

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}