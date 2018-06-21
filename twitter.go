package commbot

import (
	"context"
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// TwitterHandler wraps the function for sending out Tweets
func TwitterHandler(ctx context.Context) interface{} {
	if isTest(ctx) {
		return func() {
			testTweet(extractEnvironmentVariables(ctx))
		}
	}

	return func() {
		tweet(extractEnvironmentVariables(ctx))
	}
}

func tweet(accessToken, accessSecret, consumerKey, consumerSecret string) {
	anaconda.NewTwitterApiWithCredentials(accessToken, accessSecret, consumerKey, consumerSecret)
}

// I just want to verify that I can retrieve my account's first tweet https://twitter.com/MyReadingFeed/status/1009480446080634880
func testTweet(accessToken, accessSecret, consumerKey, consumerSecret string) string {
	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessSecret, consumerKey, consumerSecret)
	user, err := api.GetUsersShow("MyReadingFeed", url.Values{})
	if err != nil {
		log.Fatalln(err)
	}

	return user.Email
}

func extractEnvironmentVariables(ctx context.Context) (string, string, string, string) {
	at := ctx.Value(LambdaContextKey("TWITTER_ACCESS_TOKEN"))
	as := ctx.Value(LambdaContextKey("TWITTER_ACCESS_SECRET"))
	ck := ctx.Value(LambdaContextKey("TWITTER_CONSUMER_KEY"))
	cs := ctx.Value(LambdaContextKey("TWITTER_CONSUMER_SECRET"))

	return (at).(string), (as).(string), (ck).(string), (cs).(string)
}

func isTest(ctx context.Context) bool {
	test := ctx.Value(LambdaContextKey("TEST_MODE"))

	return (test).(bool)
}