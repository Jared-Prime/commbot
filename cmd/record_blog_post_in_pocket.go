package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/jared-prime/commbot"
)

var (
	enVars = []string{
		"POCKET_CONSUMER_KEY",
		"POCKET_ACCESS_TOKEN",
		"POCKET_RSS_URL",
		"BLOG_RSS_URL",
	}
)

func main() {
	ctx, err := commbot.Setup(context.Background(), enVars)
	if err != nil {
		log.Fatal(err)
	}

	handler := func() {
		for _, link := range commbot.BlogRecentLinks(ctx) {
			itemID, err := commbot.AddPocketLink(ctx, link)
			if err != nil {
				log.Fatal(err)
			}

			itemID, err = commbot.ArchivePocketLink(ctx, itemID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	lambda.Start(handler)
}
