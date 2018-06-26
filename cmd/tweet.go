package cmd

import (
	"log"

	pocket "github.com/jared-prime/commbot/pocket"
	twitter "github.com/jared-prime/commbot/twitter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tweetCmd)
}

var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Tweet from the Pocket feed.",
	Long:  "Tweet from the Pocket feed.",
	Run: func(cmd *cobra.Command, args []string) {
		pocketHandler()
	},
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
