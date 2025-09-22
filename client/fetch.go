package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchItems(client *http.Client, url string) ([]Item, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header = Headers
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return response.Items, nil
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

func PrintItems(items []Item) {
	for _, item := range items {
		//fmt.Printf("ID: %d\n", item.ID)
		fmt.Printf("Title: %s\n", item.Title)
		//fmt.Printf("URL: %s\n", item.URL)
		fmt.Printf("Status: %s\n", item.Status)
		fmt.Printf("Size: %s\n", item.SizeTitle)
		//fmt.Printf("User: %s\n", item.User.Login)
		//fmt.Printf("Photo: %s\n", item.Photo.URL)
		fmt.Printf("Price: %s %s\n", item.Price.Amount, item.Price.CurrencyCode)
		fmt.Println("--------------------------------------------------")
	}
}
