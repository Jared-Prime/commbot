package commbot

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type PocketRequest struct {
	ConsumerKey string         `json:"consumer_key"`
	AccessToken string         `json:"access_token"`
	URL         string         `json:"url,omitempty"`
	Actions     []PocketAction `json:"actions,omitempty"`
}

type PocketAction struct {
	Action string `json:"action"`
	ItemID int    `json:"item_id"`
}

type PocketResponse struct {
	Item struct {
		ItemID int `json:"item_id"`
	} `json:"item,omitempty"`
	ActionResults []bool `json:"action_results,omitempty"`
}

// ArchivePocketLink archives the given Pocket item
func ArchivePocketLink(ctx context.Context, itemID int) (int, error) {
	reqBody := buildBasePocketRequest(ctx)
	reqBody.Actions = []PocketAction{
		PocketAction{Action: "archive", ItemID: itemID},
	}
	reqBodyReader, err := json.Marshal(reqBody)
	if err != nil {
		return itemID, err
	}

	response, err := http.Post("https://getpocket.com/v3/modify", "application/json", bytes.NewBuffer(reqBodyReader))

	var pocketResponse PocketResponse
	json.NewDecoder(response.Body).Decode(&pocketResponse)

	return pocketResponse.Item.ItemID, err
}

// AddPocketLink adds the given url string to the Pocket app
func AddPocketLink(ctx context.Context, url string) (int, error) {
	reqBody := buildBasePocketRequest(ctx)
	reqBody.URL = url
	reqBodyReader, err := json.Marshal(reqBody)
	if err != nil {
		return 0, err
	}

	response, err := http.Post("https://getpocket.com/v3/add", "application/json", bytes.NewBuffer(reqBodyReader))

	var pocketResponse PocketResponse
	json.NewDecoder(response.Body).Decode(&pocketResponse)

	return pocketResponse.Item.ItemID, err
}

func buildBasePocketRequest(ctx context.Context) *PocketRequest {
	ck := ctx.Value(EnvarContextKey("POCKET_CONSUMER_KEY"))
	at := ctx.Value(EnvarContextKey("POCKET_ACCESS_TOKEN"))

	return &PocketRequest{
		ConsumerKey: (ck).(string),
		AccessToken: (at).(string),
	}
}
