package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jared-prime/commbot"
)

var (
	enVars = []string{
		"TWITTER_ACCESS_TOKEN",
		"TWITTER_ACCESS_SECRET",
		"TWITTER_CONSUMER_KEY",
		"TWITTER_CONSUMER_SECRET",
		"POCKET_RSS_URL",
		"BLOG_RSS_URL",
	}
)

func main() {
	ctx, err := commbot.Setup(context.Background(), enVars)
	if err != nil {
		log.Fatal(err)
	}

	lambda.Start(commbot.TwitterHandler(ctx))
}
