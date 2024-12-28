package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create videos table if it doesn't exist
	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS videos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		video_id VARCHAR(50) NOT NULL UNIQUE,
		title TEXT NOT NULL,
		description TEXT,
		channel_title VARCHAR(255),
		published_at DATETIME,
		thumbnail_url VARCHAR(255),
		live_broadcast VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Database tables initialized")
	return nil
}