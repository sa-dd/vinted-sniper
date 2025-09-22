package main

import (
	"fmt"
	"net/http"
	"time"
	"vinted-sniper/client"
)

const url = "https://www.vinted.co.uk/api/v2/catalog/items?page=1&per_page=96&time=1758397863&global_search_session_id=d9bc1e4f-6b8d-45a1-90d8-fb8a041ef637&search_text=uggs+slippers&catalog_ids=&order=newest_first&size_ids=&brand_ids=&status_ids=&color_ids=&material_ids="

func main() {

	httpClient := &http.Client{}
	var latestItemId int

	for {
		items, err := client.FetchItems(httpClient, url)
		if err != nil {
			fmt.Printf("Error fetching items: %v\n", err)
		} 

		if latestItemId != 0 {
			latestItems := client.FindLatestItems(latestItemId, items)
			fmt.Printf("------Latest Items-----------")
			client.PrintItems(latestItems)
		}
		if len(items) > 0 {
			latestItemId = items[0].Id
		}
		}
		// Wait for 10 seconds before next fetch
		time.Sleep(60 * time.Second)
}
}
