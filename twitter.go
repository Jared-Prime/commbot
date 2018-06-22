package commbot

import (
	"context"
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// TwitterHandler wraps the function for sending out Tweets
func TwitterHandler(ctx context.Context) func() {
	return func() {
		tweetPocketLinks(ctx)
	}
}

func tweetPocketLinks(ctx context.Context) {
	accessToken, accessSecret, consumerKey, consumerSecret := extractEnvironmentVariables(ctx)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	for _, link := range PocketRecentLinks(ctx) {
		_, err := api.PostTweet(link, url.Values{})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func extractEnvironmentVariables(ctx context.Context) (string, string, string, string) {
	at := ctx.Value(EnvarContextKey("TWITTER_ACCESS_TOKEN"))
	as := ctx.Value(EnvarContextKey("TWITTER_ACCESS_SECRET"))
	ck := ctx.Value(EnvarContextKey("TWITTER_CONSUMER_KEY"))
	cs := ctx.Value(EnvarContextKey("TWITTER_CONSUMER_SECRET"))

	return (at).(string), (as).(string), (ck).(string), (cs).(string)
}
