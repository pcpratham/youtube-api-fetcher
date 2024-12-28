package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
	"github.com/pcpratham/youtube-api-fetcher/internal/config"
	"github.com/pcpratham/youtube-api-fetcher/internal/database"
	"github.com/pcpratham/youtube-api-fetcher/internal/handlers"
	"github.com/pcpratham/youtube-api-fetcher/internal/service"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	cfg := config.Load()
	fmt.Println("cfg :- ",cfg.DatabaseURL)

	
	db, err := database.InitDB(cfg.DatabaseURL)
	fmt.Println("cfg :- ",os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	videoService := service.NewYouTubeService(cfg.YouTubeAPIKey)
	videoHandler := handlers.NewVideoHandler(db, videoService)

	c := cron.New()
	c.AddFunc("@every 10s", func() { 
		videoService.FetchAndSaveVideos(db, "food") 
	})
	c.Start()
	defer c.Stop()

	// Setting up routes
	http.HandleFunc("/videos", videoHandler.GetVideos)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}