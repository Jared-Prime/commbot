package cmd

import (
	"log"

	pocket "github.com/jared-prime/commbot/pocket"
	twitter "github.com/jared-prime/commbot/twitter"
	"github.com/spf13/cobra"
)

var (
	from []string
)

func init() {
	rootCmd.AddCommand(tweetCmd)

	tweetCmd.Flags().StringArrayVarP(&from, "from", "i", []string{"STDOUT"}, "sources for tweeting messages")
}

var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Tweet from the given source.",
	Long:  "Tweet from the given source.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, handler := range makeHandlers(args) {
			handler()
		}
	},
}

func makeHandlers(args []string) []func() {
	var handlers []func()

	for _, source := range from {
		switch source {
		case "pocket":
			handlers = append(handlers, pocketHandler)
		case "STDOUT":
			handlers = append(handlers, twitterHandler(args))
		default:
			log.Fatalln("unknown source for tweets: ", from)
		}
	}

	return handlers
}

func twitterHandler(messages []string) func() {
	return func() {
		err := twitter.TweetMessages(accessToken, accessSecret, consumerKey, consumerSecret, messages)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func pocketHandler() {
	var links []string
	items, err := pocket.Recent(pocketRSSURL)
	if err != nil {
		log.Println("unable to retrieve Pocket feed: ", pocketRSSURL)
		log.Println(err)
		log.Println("skipping...")
	}

	for _, item := range items {
		links = append(links, item.Link)
	}

	tweet := twitterHandler(links)
	tweet()
}
