package service

import (
	"database/sql"
	"log"

	"github.com/pcpratham/youtube-api-fetcher/internal/models"
	"github.com/pcpratham/youtube-api-fetcher/pkg/youtube"
)

type YouTubeService struct {
	client *youtube.Client
}

func NewYouTubeService(apiKey string) *YouTubeService {
	return &YouTubeService{
		client: youtube.NewClient(apiKey),
	}
}

func (s *YouTubeService) FetchAndSaveVideos(db *sql.DB, searchQuery string) error {
	log.Printf("Starting to fetch videos for query: %s", searchQuery)
	
	videos, err := s.client.FetchVideos(searchQuery)
	if err != nil {
		log.Printf("Error fetching videos: %v", err)
		return err
	}

	log.Printf("Successfully fetched %d videos, preparing to save", len(videos))

	err = s.saveVideosToDB(db, videos)
	if err != nil {
		log.Printf("Error saving videos: %v", err)
		return err
	}

	return nil
}

func (s *YouTubeService) saveVideosToDB(db *sql.DB, videos []models.Video) error {
	successCount := 0
	
	for _, video := range videos {
		query := `
			INSERT IGNORE INTO videos (
				video_id, title, description, channel_title, 
				published_at, thumbnail_url, live_broadcast
			)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`
		result, err := db.Exec(
			query,
			video.VideoID,
			video.Title,
			video.Description,
			video.ChannelTitle,
			video.PublishedAt,
			video.ThumbnailURL,
			video.LiveBroadcast,
		)
		if err != nil {
			log.Printf("Failed to save video %s: %v", video.VideoID, err)
			continue
		}

		affected, err := result.RowsAffected()
		if err != nil {
			log.Printf("Error getting rows affected for video %s: %v", video.VideoID, err)
			continue
		}

		if affected > 0 {
			successCount++
			log.Printf("Successfully saved video: %s - %s", video.VideoID, video.Title)
		} else {
			log.Printf("Video already exists in database: %s", video.VideoID)
        }
	}
	
	log.Printf("Saved %d new videos to database (out of %d total videos fetched)", successCount, len(videos))
	return nil
}