# Pro URL Shortener (Go + SQLite)

A professional-grade URL shortening service built with Go. This version features permanent data storage using SQLite, click analytics, and a JSON-based API.

## 🚀 Key Features

- **Persistence:** Uses SQLite to ensure your shortened links survive server restarts.
- **Click Tracking:** Automatically tracks how many times each shortened link is visited.
- **JSON API:** Follows modern API standards by accepting and returning JSON payloads.
- **Concurrency Safe:** Designed to handle multiple requests reliably using Go's standard library.

## 🛠️ Prerequisites

To run this project, you need:
- **Go** installed on your system.
- A C compiler (like GCC) if using the `github.com/mattn/go-sqlite3` driver, OR you can switch to a pure Go driver like `modernc.org/sqlite`.

## 📦 Installation

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/YOUR_USERNAME/go-url-shortener.git](https://github.com/YOUR_USERNAME/go-url-shortener.git)
   cd go-url-shortener