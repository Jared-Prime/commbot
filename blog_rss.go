package commbot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// BlogRecentLinks returns the urls to most recent pocket items
func BlogRecentLinks(ctx context.Context) []string {
	return recentBlogLinks(GetBlogFeed(ctx))
}

// GetBlogFeed parses the RSS feed associated with the given context
func GetBlogFeed(ctx context.Context) *gofeed.Feed {
	gf := gofeed.NewParser()
	feed, err := gf.ParseURL(extractBlogRssURL(ctx))
	if err != nil {
		log.Fatalln(err)
	}

	return feed
}

func recentBlogLinks(feed *gofeed.Feed) []string {
	var items []string

	for _, item := range feed.Items {
		if time.Since(*item.PublishedParsed) < time.Hour*4 {
			items = append(items, buildSummary(item))
		}
	}

	return items
}

func buildSummary(item *gofeed.Item) string {
	return fmt.Sprintf("%s\n%s", item.Link, item.Description)
}

func extractBlogRssURL(ctx context.Context) string {
	url := ctx.Value(EnvarContextKey("BLOG_RSS_URL"))

	return (url).(string)
}
