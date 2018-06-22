package main

import (
	"context"
	"log"

	"github.com/jared-prime/commbot"
)

func main() {
	ctx, err := commbot.Setup(context.Background(), []string{"POCKET_RSS_URL", "BLOG_RSS_URL"})
	if err != nil {
		log.Fatalln(err)
	}

	links := commbot.BlogRecentLinks(ctx)
	links = append(links, commbot.PocketRecentLinks(ctx)...)

	log.Println(len(links))

	for _, link := range links {
		log.Println(link)
	}
}
