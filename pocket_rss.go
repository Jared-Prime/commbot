package commbot

import (
	"context"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// PocketRecentLinks returns the urls to most recent pocket items
func PocketRecentLinks(ctx context.Context) []string {
	return recentLinks(GetPocketFeed(ctx))
}

// GetPocketFeed parses the RSS feed associated with the given context
func GetPocketFeed(ctx context.Context) *gofeed.Feed {
	gf := gofeed.NewParser()
	feed, err := gf.ParseURL(extractRssURL(ctx))
	if err != nil {
		log.Fatalln(err)
	}

	return feed
}

func recentLinks(feed *gofeed.Feed) []string {
	var items []string

	for _, item := range feed.Items {
		if time.Since(*item.PublishedParsed) < time.Hour*24 {
			items = append(items, item.Link)
		}
	}

	return items
}

func extractRssURL(ctx context.Context) string {
	url := ctx.Value(EnvarContextKey("POCKET_RSS_URL"))

	return (url).(string)
}
