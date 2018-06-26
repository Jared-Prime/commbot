package commbot

import (
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// TweetMessages sends out the messages as indiviual tweets
func TweetMessages(accessToken, accessSecret, consumerKey, consumerSecret string, messages []string) error {
	api := newAPI(accessToken, accessSecret, consumerKey, consumerSecret)

	for _, message := range messages {
		_, err := api.PostTweet(message, url.Values{})
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteOldMessages remove all Twitter messages older than the given number of hours
func DeleteOldMessages(accessToken, accessSecret, consumerKey, consumerSecret, since string) error {
	api := newAPI(accessToken, accessSecret, consumerKey, consumerSecret)

	timeline, err := timeline(api)
	if err != nil {
		return err
	}

	maxAge, err := time.ParseDuration(since)
	if err != nil {
		return err
	}

	for _, tweet := range timeline {
		createdAt, err := tweet.CreatedAtTime()
		if err != nil {
			return err
		}

		if time.Since(createdAt) > maxAge {
			if _, err := api.DeleteTweet(tweet.Id, true); err != nil {
				return err
			}
		}
	}

	return nil
}

func timeline(api *anaconda.TwitterApi) ([]anaconda.Tweet, error) {
	args := url.Values{}
	args.Add("count", "200")
	args.Add("include_rts", "true")

	return api.GetUserTimeline(args)
}

func newAPI(accessToken, accessSecret, consumerKey, consumerSecret string) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	return anaconda.NewTwitterApi(accessToken, accessSecret)
}
