package config

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
)

type Config struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
}

var TwitterClient = &oauth1.Config{
	ConsumerKey:    os.Getenv("TWITTER_API_KEY"),
	ConsumerSecret: os.Getenv("TWITTER_API_SECRET"),
	CallbackURL:    "http://127.0.0.1:8080/twitter/callback",
	Endpoint:       twitterOAuth1.AuthorizeEndpoint,
}

func GetClient() *twitter.Client {
	config := TwitterClient

	// takes access token and access token secret.
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_TOKEN_SECRET"))

	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}

// using OAuth2 Client.
// ID: bTJ2eEZoRlVMWF9yUVI1d3I0UWo6MTpjaQ
// Secrete: YK0HDsmD_E4DB0s3Sd8B77IDnQT3hKM6vhvCwi_sUfpGxIQ-E2
