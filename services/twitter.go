package services

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	config2 "twitterSearchSentimentalAnalysis/config"
	"twitterSearchSentimentalAnalysis/entity"
	"twitterSearchSentimentalAnalysis/logging"
	"twitterSearchSentimentalAnalysis/repository"
)

func DoTwitterSearch(term string) {
	client := getTwitterClient()

	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: term,
		Lang:  "pt",
		Count: 1,
	})
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets.Statuses {
		go processAndSaveTweet(tweet)
	}
}

func processAndSaveTweet(tweet twitter.Tweet) {
	sentiment, score := GetSentimentalAnalysis(tweet.Text)
	filteredTweet := entity.FilteredTweet{
		IdTweet:        tweet.ID,
		Tweet:          tweet.Text,
		Author:         tweet.User.Name,
		SentimentScore: score,
		SentimentType:  sentiment,
	}
	repository.SaveTweet(filteredTweet)
	logging.LogToFile(fmt.Sprintf("%s -> %s -> %f", tweet.Text, sentiment, *score), "default")
}

func getTwitterClient() *twitter.Client {
	config := oauth1.NewConfig(config2.GetTwitterApiKey(), config2.GetTwitterSecretKey())
	token := oauth1.NewToken(config2.GetTwitterAccessToken(), config2.GetTwitterAccessSecret())
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return twitter.NewClient(httpClient)

}
