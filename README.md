# YouTube API Fetcher

A Go-based project that fetches the latest YouTube videos for a predefined search query using the YouTube Data API, stores them in a MySQL database, and provides an API for paginated responses. This project supports multiple API keys for rate limit handling and includes a background fetch mechanism for keeping data up-to-date.

## Features

- Fetch YouTube videos based on a predefined search query.
- Store video data (title, description, etc.) in a MySQL database.
- Provide paginated responses for easy client-side consumption.
- Background job to fetch and update video data periodically.
- Support for multiple YouTube API keys for better rate limit handling.
- Optional dashboard integration for monitoring and management.

---

## Prerequisites

- **Go 1.20+**
- **MySQL 8+**
- **YouTube Data API Key(s)** (from the Google Cloud Console)

---

## Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/pcpratham/youtube-api-fetcher.git
   cd youtube-api-fetcher
2. **Creation of .env file:**
```bash
   YOUTUBE_API_KEY=your_youtube_api_key
   DATABASE_URL=your_database_url
   PORT=your_port_number
```
3. **Run to install all dependencies:**
```bash
   go mod tidy
```
4. **Set up database:**
```bash
   CREATE DATABASE youtube_fetcher;
```
5. **Run command:**
```bash
    go run ./cmd/server/main.go
```
