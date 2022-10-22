package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"tweet/tweedy/auth"
	"tweet/tweedy/config"
)

func main() {

	const addr = "localhost:8080"

	// read credentials from environment variables.
	creds := &config.Config{
		TwitterConsumerKey:    os.Getenv("TWITTER_API_KEY"),
		TwitterConsumerSecret: os.Getenv("TWITTER_API_SECRET"),
	}

	// allow consumer credential flag to overide defaults.
	consumerKey := flag.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flag.String("consumer-secret", "", "Twitter Consumer Secret")
	flag.Parse()

	if *consumerKey != "" {
		creds.TwitterConsumerKey = *consumerKey
	}

	if *consumerSecret != "" {
		creds.TwitterConsumerSecret = *consumerSecret
	}

	if creds.TwitterConsumerKey == "" {
		log.Fatal("Missing Twitter Consumer Key")
	}

	if creds.TwitterConsumerSecret == "" {
		log.Fatal("Missing Twitter Consumer Secret")
	}

	log.Printf("Starting Tweedy Server on %s\n", addr)

	err := http.ListenAndServe(addr, auth.New(creds))

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Fatal("Tweedy server active")
}
