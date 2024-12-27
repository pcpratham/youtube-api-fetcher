package scheduler

import (
	"os"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
)

type Video struct {
	Title       string `json:"title"`
    Description string `json:"description"`
    PublishTime string `json:"publishTime"`
    Thumbnail   string `json:"thumbnail"`
    VideoID     string `json:"videoId"`
}


var apiKeyIndex int

func getAPIKeys() string {
	keys := strings.Split(os.Getenv("YOUTUBE_API_KEYS"), ",")
	key := keys[apiKeyIndex]
	apiKeyIndex = (apiKeyIndex + 1) % len(keys)
	return key 
}


func FetchYoutubeData(query string) ([]Video,error){
	apiKey := getAPIKeys()
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=10&q=%s&key=%s", query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Print("Error in fetching data from youtube via API")
		return nil, err	
	}

	defer resp.Body.Close()

	var result struct {
        Items []struct {
            Snippet struct {
                Title       string `json:"title"`
                Description string `json:"description"`
                PublishTime string `json:"publishTime"`
                Thumbnails  struct {
                    Default struct {
                        URL string `json:"url"`
                    } `json:"default"`
                } `json:"thumbnails"`
            } `json:"snippet"`
            ID struct {
                VideoID string `json:"videoId"`
            } `json:"id"`
        } `json:"items"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    videos := make([]Video, 0)
    for _, item := range result.Items {
        videos = append(videos, Video{
            Title:       item.Snippet.Title,
            Description: item.Snippet.Description,
            PublishTime: item.Snippet.PublishTime,
            Thumbnail:   item.Snippet.Thumbnails.Default.URL,
            VideoID:     item.ID.VideoID,
        })
    }

    return videos, nil
}