package cmd

import (
	"log"

	"github.com/jared-prime/commbot/pocket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(recentCmd)
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
