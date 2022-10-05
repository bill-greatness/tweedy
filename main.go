package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

func main() {
	fmt.Print("Welcome to tweedy!")
	
	// create router.
	router := http.NewServeMux()

}

// Get User Authenticated.
func getClient() *twitter.Client{

	config := &clientcredentials.Config{
		ClientID:"Some ID Here", 
		ClientSecret: "Some Secret Here",
		TokenURL:"https://api.twitter.com/oauth2/token",
	}

	client := config.Client(oauth2.NoContext)
	
	return twitter.NewClient(client)

}

func getTweets(){

}