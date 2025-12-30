# Pro URL Shortener (Go + SQLite)

A professional-grade URL shortening service built with Go. This version features permanent data storage using SQLite, click analytics, and a JSON-based API.

## Key Features

- Persistence: Uses SQLite to ensure your shortened links survive server restarts.
- Click Tracking: Automatically tracks how many times each shortened link is visited.
- JSON API: Follows modern API standards by accepting and returning JSON payloads.
- Concurrency Safe: Designed to handle multiple requests reliably using Go's standard library.

## Project Structure

- main.go: The entry point containing the server setup, database initialization, and HTTP handlers.
- links.db: (Auto-generated) The SQLite database file storing your link data.
- go.mod: The Go module file managing dependencies.

## Prerequisites

To run this project, you need:
- Go installed on your system.
- A C compiler (like GCC) if using the github.com/mattn/go-sqlite3 driver.

## Installation

1. Clone the repository:
   git clone https://github.com/YOUR_USERNAME/go-url-shortener.git
   cd go-url-shortener

2. Install dependencies:
   go get github.com/mattn/go-sqlite3

3. Run the server:
   go run main.go

## API Usage

### 1. Shorten a URL
Endpoint: POST /shorten
Body:
{
  "url": "https://www.example.com/very/long/path"
}

Example Curl Command:
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d "{\"url\": \"https://golang.org\"}"

### 2. Redirect
Simply visit the generated short link in your browser:
http://localhost:8080/{short_code}

## Database Schema

The application automatically creates a links.db file with the following structure:
- id: Primary Key (Integer)
- original_url: The destination address (Text)
- short_code: The unique 6-character identifier (Text)
- clicks: Number of times the link was used (Integer)

## Error Handling

The API includes basic error handling for the following scenarios:
- 405 Method Not Allowed: If a user tries to GET the /shorten endpoint.
- 400 Bad Request: If the JSON body is malformed or the URL parameter is missing.
- 404 Not Found: If the short code provided does not exist in the database.
- 500 Internal Server Error: If there is an issue writing to or reading from the database.

## Future Improvements

- Add a GUI: Build a simple HTML/JavaScript frontend to interact with the API.
- Custom Aliases: Allow users to specify their own short codes (e.g., localhost:8080/my-link).
- Expiration Dates: Add a column to the database to automatically disable links after a certain period.
- User Accounts: Implement authentication to allow users to manage their own specific links.