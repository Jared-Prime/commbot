package cmd

import (
	"log"

	"github.com/jared-prime/commbot/pocket"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(archiveCmd)
}

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive links into Pocket.",
	Long:  "Archive links into Pocket.",
	Run: func(cmd *cobra.Command, linksForArchive []string) {
		for _, link := range linksForArchive {
			itemID, err := commbot.Add(link)
			if err != nil {
				log.Fatal(err)
			}

			_, err = commbot.Archive(itemID)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}
