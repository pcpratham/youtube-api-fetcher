package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

    "github.com/pcpratham/youtube-api-fetcher/internal/models"
)

type Client struct {
	apiKey        string
	lastPageToken string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) FetchVideos(searchQuery string) ([]models.Video, error) {
	encodedQuery := url.QueryEscape(searchQuery)
	
	baseURL := "https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=10&q=%s&type=video&key=%s"
	apiURL := fmt.Sprintf(baseURL, encodedQuery, c.apiKey)
	
	if c.lastPageToken != "" {
		apiURL += "&pageToken=" + c.lastPageToken
	}

	log.Printf("Fetching videos from URL: %s", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from YouTube API: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if the response status is not 200
	if resp.StatusCode != http.StatusOK {
		log.Printf("YouTube API Error Response: %s", string(body))
		return nil, fmt.Errorf("YouTube API returned status code %d", resp.StatusCode)
	}

	var result struct {
		NextPageToken string `json:"nextPageToken"`
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title             string `json:"title"`
				Description       string `json:"description"`
				ChannelTitle      string `json:"channelTitle"`
				PublishedAt       string `json:"publishedAt"`
				Thumbnails        struct {
					High struct {
						URL string `json:"url"`
					} `json:"high"`
				} `json:"thumbnails"`
				LiveBroadcastContent string `json:"liveBroadcastContent"`
			} `json:"snippet"`
		} `json:"items"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Response body: %s", string(body))
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	log.Printf("Received %d items from YouTube API", len(result.Items))

	c.lastPageToken = result.NextPageToken

	var videos []models.Video
	for _, item := range result.Items {
		video := models.Video{
			VideoID:       item.ID.VideoID,
			Title:         item.Snippet.Title,
			Description:   item.Snippet.Description,
			ChannelTitle:  item.Snippet.ChannelTitle,
			PublishedAt:   item.Snippet.PublishedAt,
			ThumbnailURL:  item.Snippet.Thumbnails.High.URL,
			LiveBroadcast: item.Snippet.LiveBroadcastContent,
		}
		videos = append(videos, video)
		log.Printf("Processing video: ID=%s, Title=%s", video.VideoID, video.Title)
	}

	return videos, nil
}