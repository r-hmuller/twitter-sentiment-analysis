package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"twitterSearchSentimentalAnalysis/config"
	"twitterSearchSentimentalAnalysis/logging"
)

func GetSentimentalAnalysis(tweet string) (string, *float32) {
	client := config.GetAwsComprehendClient()
	sentiment, err := client.DetectSentiment(context.TODO(),
		&comprehend.DetectSentimentInput{Text: &tweet, LanguageCode: "pt"})
	if err != nil {
		logging.LogToFile(err.Error(), "default")
		return "", aws.Float32(0.0)
	}

	resultantSentiment := string(sentiment.Sentiment)
	var resultantScore *float32
	switch resultantSentiment {
	case "NEGATIVE":
		resultantScore = sentiment.SentimentScore.Negative
	case "POSITIVE":
		resultantScore = sentiment.SentimentScore.Positive
	case "NEUTRAL":
		resultantScore = sentiment.SentimentScore.Neutral
	case "MIXED":
		resultantScore = sentiment.SentimentScore.Mixed
	default:
		resultantScore = aws.Float32(0.0)
	}

	return resultantSentiment, resultantScore
}
