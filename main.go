package main

import (
	"os"
	"regexp"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig(os.Getenv("consumer_key"), os.Getenv("consumer_secret"))
	token := oauth1.NewToken(os.Getenv("access_token"), os.Getenv("access_token_secret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	var sinceID int64 = 0
	var mentionTimelineParams *twitter.MentionTimelineParams = &twitter.MentionTimelineParams{}

	for {
		if sinceID != 0 {
			mentionTimelineParams = &twitter.MentionTimelineParams{SinceID: sinceID}
		}

		mentions, _, _ := client.Timelines.MentionTimeline(mentionTimelineParams)

		if len(mentions) != 0 {
			sinceID = mentions[0].ID
			for _, mention := range mentions {
				var userTweet, _, _ = client.Statuses.Show(mention.InReplyToStatusID, nil)

				var userPattern = regexp.MustCompile(`(@[a-zA-Z0-9_]{0,}\s)`)
				var tweet = userPattern.ReplaceAllString(userTweet.Text, "")
				var vowelPattern = regexp.MustCompile(`[aiueo]`)
				var replaceTweet = vowelPattern.ReplaceAllString(tweet, "i")

				statusUpdateParams := &twitter.StatusUpdateParams{InReplyToStatusID: mention.ID}
				_, _, _ = client.Statuses.Update("@"+mention.User.ScreenName+" "+replaceTweet, statusUpdateParams)
			}
		}
		time.Sleep(30 * time.Second)
	}
}
