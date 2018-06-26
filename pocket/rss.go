package commbot

import (
	"time"

	"github.com/mmcdole/gofeed"
)

// Recent returns a collection of the most recent Pocket items for the given feed
func Recent(url string) ([]*gofeed.Item, error) {
	var items []*gofeed.Item

	feed, err := GetFeed(url)
	if err != nil {
		return items, err
	}

	for _, item := range feed.Items {
		if time.Since(*item.PublishedParsed) > time.Hour*4 {
			items = append(items, item)
		}
	}

	return items, err
}

// GetFeed returns a Pocket RSS feed object for the given URL
func GetFeed(url string) (*gofeed.Feed, error) {
	parser := gofeed.NewParser()
	return parser.ParseURL(url)
}
