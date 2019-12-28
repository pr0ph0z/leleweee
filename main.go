package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig(os.Getenv("consumer_key"), os.Getenv("consumer_secret"))
	token := oauth1.NewToken(os.Getenv("access_token"), os.Getenv("access_token_secret"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	tweet, resp, err := client.Statuses.Update("Initial commit! Yeay", nil)
	if tweet != nil {
		fmt.Println(tweet)
	}
	if resp != nil {
		fmt.Println(resp)
	}
	if err != nil {
		fmt.Println(err)
	}
}
