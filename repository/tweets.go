package repository

import (
	"context"
	"twitterSearchSentimentalAnalysis/config"
	"twitterSearchSentimentalAnalysis/entity"
	"twitterSearchSentimentalAnalysis/logging"
)

func SaveTweet(tweet entity.FilteredTweet) {
	conn := config.GetPostgresConnection()
	if _, err := conn.Exec(context.Background(), "INSERT INTO "+
		"tweets(idtweet, tweet, author, sentimenttype, sentimentscore) VALUES($1, $2, $3, $4, $5)",
		tweet.IdTweet, tweet.Tweet, tweet.Author, tweet.SentimentType, tweet.SentimentScore); err != nil {
		// Handling error, if occur
		logging.LogToFile("Erro -> "+err.Error(), "default")
		return
	}
}
