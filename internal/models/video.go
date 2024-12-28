package models

type Video struct {
	VideoID       string `json:"videoId"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ChannelTitle  string `json:"channelTitle"`
	PublishedAt   string `json:"publishedAt"`
	ThumbnailURL  string `json:"thumbnailUrl"`
	LiveBroadcast string `json:"liveBroadcastContent"`
}