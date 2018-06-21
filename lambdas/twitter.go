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
	}
)

func main() {
	ctx, err := setup(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	lambda.Start(commbot.TwitterHandler(ctx))
}

func setup(ctx context.Context) (context.Context, error) {
	type LambdaContextKey string

	_, err := commbot.Envar("LIVE_MODE")
	if err != nil {
		// set test node to true
		ctx = context.WithValue(ctx, LambdaContextKey("TEST_MODE"), true)
		ctx = context.WithValue(ctx, LambdaContextKey("LIVE_MODE"), false)
	}

	for _, name := range enVars {
		value, err := commbot.Envar(name)

		if err != nil {
			return ctx, err
		}

		ctx = context.WithValue(ctx, LambdaContextKey(name), value)
	}

	return ctx, nil
}
