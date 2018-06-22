package commbot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// PocketRecentLinks returns the urls to most recent pocket items
func PocketRecentLinks(ctx context.Context) []string {
	return recentPocketLinks(GetPocketFeed(ctx))
}

// GetPocketFeed parses the RSS feed associated with the given context
func GetPocketFeed(ctx context.Context) *gofeed.Feed {
	gf := gofeed.NewParser()
	feed, err := gf.ParseURL(extractPocketRssURL(ctx))
	if err != nil {
		log.Fatalln(err)
	}

	return feed
}

func recentPocketLinks(feed *gofeed.Feed) []string {
	var items []string

	for _, item := range feed.Items {
		if time.Since(*item.PublishedParsed) < time.Hour*24 {
			items = append(items, buildLink(item))
		}
	}

	return items
}

func buildLink(item *gofeed.Item) string {
	return fmt.Sprintf("%s\n%s", item.Title, item.Link)
}

func extractPocketRssURL(ctx context.Context) string {
	url := ctx.Value(EnvarContextKey("POCKET_RSS_URL"))

	return (url).(string)
}
