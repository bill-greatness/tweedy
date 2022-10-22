package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tweet/tweedy/config"

	"github.com/dghubble/go-twitter/twitter"
)

var client = config.GetClient()

func GetTimelineTweets(w http.ResponseWriter, r *http.Request) {

	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Print(tweets, resp.StatusCode)

}

func GetUserTweets(w http.ResponseWriter, req *http.Request) {
	tweets, resp, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count:   5,
		SinceID: 4500,
	})

	var Tweets []*twitter.Tweet
	var UserTweetInfo []*UserTweets
	// return error if True
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Print(resp.Body)

	for _, tweet := range tweets {
		Tweets = append(Tweets, &tweet)
	}

	info, err := json.Marshal(Tweets)
	fmt.Print(string(info))
	// decode information to UserInfo Tweets.
	if err := json.Unmarshal(info, &UserTweetInfo); err != nil {
		http.Error(w, "Something Went Wrong", http.StatusBadRequest)
	}

	b, err := json.MarshalIndent(&UserTweetInfo, " ", " ")

	fmt.Fprintf(w, "%v", string(b))

}
