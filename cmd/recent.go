package cmd

import (
	"log"
	"os"

	"github.com/jared-prime/commbot/pocket"
	"github.com/spf13/cobra"
)

var (
	pocketRSSURL string
)

func init() {
	rootCmd.AddCommand(recentCmd)
	recentCmd.Flags().StringVarP(&pocketRSSURL, "feed", "f", os.Getenv("POCKET_RSS_URL"), "URL for the Pocket RSS feed")
}

var recentCmd = &cobra.Command{
	Use:   "recent",
	Short: "List recently read articles in Pocket.",
	Long:  "List recently read articles in Pocket.",
	Run: func(cmd *cobra.Command, args []string) {
		items, err := commbot.Recent(pocketRSSURL)
		if err != nil {
			log.Fatalln(err)
		}

		for _, item := range items {
			log.Println(item.Title, ": ", item.Link)
		}
	},
}
