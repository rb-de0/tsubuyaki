package main

import (
	"fmt"
	"os"
)

func main() {

	router := NewRouter(os.Args)
	router.AddHandler("tweet", tweet)
	router.Execute()
}

func tweet(args []string) {
	if len(args) <= 2 {
		fmt.Println("Error, no tweet message")
		os.Exit(-1)
	}

	tweetMessage := args[2]

	api := NewApi()
	api.PostTweet(tweetMessage, nil)
}
