package config

import (
	"context"
	_ "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"twitterSearchSentimentalAnalysis/logging"
)

func GetTwitterApiKey() string {
	return os.Getenv("TWITTER_API_KEY")
}

func GetTwitterSecretKey() string {
	return os.Getenv("TWITTER_SECRET_KEY")
}

func GetTwitterAccessToken() string {
	return os.Getenv("TWITTER_ACCESS_TOKEN")
}

func GetTwitterAccessSecret() string {
	return os.Getenv("TWITTER_ACCESS_SECRET")
}

func GetAwsComprehendClient() *comprehend.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return comprehend.NewFromConfig(cfg)
}

func GetPostgresConnection() *pgxpool.Pool {
	urlConnection := os.Getenv("POSTGRES_URL")
	conn, err := pgxpool.Connect(context.Background(), urlConnection)
	if err != nil {
		logging.LogToFile(err.Error(), "fatal")
	}
	return conn
}
