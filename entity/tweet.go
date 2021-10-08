package entity

import "time"

type FilteredTweet struct {
	ID             int
	IdTweet        int64
	Tweet          string
	Author         string
	SentimentType  string
	SentimentScore *float32
	CreatedAt      time.Time
}
