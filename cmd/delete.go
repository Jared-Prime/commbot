package cmd

import (
	"log"

	twitter "github.com/jared-prime/commbot/twitter"
	"github.com/spf13/cobra"
)

var hoursSince string

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&hoursSince, "since", "s", "36h", "max age of Tweet")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tweets.",
	Long:  "Delete tweets.",
	Run: func(cmd *cobra.Command, args []string) {
		err := twitter.DeleteOldMessages(accessToken, accessSecret, consumerKey, consumerSecret, hoursSince)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
