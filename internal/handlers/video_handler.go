package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pcpratham/youtube-api-fetcher/internal/models"
	"github.com/pcpratham/youtube-api-fetcher/internal/service"
)

type VideoHandler struct {
	db      *sql.DB
	service *service.YouTubeService
}

func NewVideoHandler(db *sql.DB, service *service.YouTubeService) *VideoHandler {
	return &VideoHandler{
		db:      db,
		service: service,
	}
}

func (h *VideoHandler) GetVideos(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, video_id, title, description, channel_title, published_at, thumbnail_url, live_broadcast, created_at 
		FROM videos 
		ORDER BY created_at DESC
		LIMIT 50
	`

	rows, err := h.db.Query(query)
	if err != nil {
		log.Printf("Database query error: %v", err)
		http.Error(w, "Failed to fetch videos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	videos := []models.Video{}
	for rows.Next() {
		var video models.Video
		var id int
		var publishedAtStr sql.NullString
		var createdAt string

		err := rows.Scan(
			&id,
			&video.VideoID,
			&video.Title,
			&video.Description,
			&video.ChannelTitle,
			&publishedAtStr,
			&video.ThumbnailURL,
			&video.LiveBroadcast,
			&createdAt,
		)
		if err != nil {
			log.Printf("Row scanning error: %v", err)
			continue
		}

		if publishedAtStr.Valid {
			video.PublishedAt = publishedAtStr.String
		}

		videos = append(videos, video)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}