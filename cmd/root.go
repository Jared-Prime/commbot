package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commbot",
	Short: "Communications robot for AWS Lambda",
	Long:  "Communications robot for AWS Lambda",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("No actions for root. Try `commbot --help`")
	},
}

// Execute defines the default command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&accessToken, "twiiter_access_token", "t", os.Getenv("TWITTER_ACCESS_TOKEN"), "Twitter access token")
	rootCmd.PersistentFlags().StringVarP(&accessSecret, "twiiter_access_secret", "s", os.Getenv("TWITTER_ACCESS_SECRET"), "Twitter access token")
	rootCmd.PersistentFlags().StringVarP(&consumerKey, "twiiter_consumer_key", "k", os.Getenv("TWITTER_CONSUMER_KEY"), "Twitter access token")
	rootCmd.PersistentFlags().StringVarP(&consumerSecret, "twiiter_consumer_secret", "c", os.Getenv("TWITTER_CONSUMER_SECRET"), "Twitter access token")
}
