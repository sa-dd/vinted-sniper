package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const base_refresh = "https://www.vinted.co.uk/web/api/auth/refresh"

func fetch_cookies(client *http.Client) error {
	req, err := http.NewRequest("POST", base_refresh, nil)
	if err != nil {
		return fmt.Errorf("create request failed for session refresh: %w", err)
	}

	req.Header = Headers
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	fmt.Println(resp.Status)

	cookies := resp.Cookies()
	AccessToken = cookies[0].Value
	RefreshToken = cookies[1].Value
	return nil
}

func FetchItems(client *http.Client, url string) ([]Item, error) {
	body, err := Get(url)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var parsedResponse Response
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return parsedResponse.Items, nil
}

func FindLatestItems(latestItemId int, items []Item) []Item {
	var latestItems []Item

	for _, item := range items {
		if item.Id == latestItemId {
			break
		}
		latestItems = append(latestItems, item)

	}

	return latestItems
}
