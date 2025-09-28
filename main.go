package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vinted-sniper/client"
	"vinted-sniper/discord"
)

const (
	url = "https://www.vinted.co.uk/api/v2/catalog/items?page=1&per_page=96&time=1758555430&global_search_session_id=8d83c2b9-8740-424c-a6c1-943c66c6f1f4&search_text=ugg%20tasman&currency=GBP&order=newest_first"
)

var (
	latestItemId int = 7143504697
)

func main() {

	httpClient := &http.Client{}
	dg, err := discord.Initiate()
	if err != nil {
		log.Println("Error initializing Bot instance: ", err)
	}

	for {
		items, err := client.FetchItems(httpClient, url)
		if err != nil {
			fmt.Printf("Error fetching items: %v\n", err)
		}

		if latestItemId != 0 {
			latestItems := client.FindLatestItems(latestItemId, items)
			fmt.Println(fmt.Sprintf("Latest item ID: %d", latestItemId))
			fmt.Println(fmt.Sprintf("Latest Items: %d", len(latestItems)))
			if len(latestItems) != 0 {
				notifs := discord.CreateEmbed(latestItems)
				discord.SendNotif(dg, notifs)
			}
		}

		if len(items) > 0 {
			latestItemId = items[0].Id
		}

		_, notifErr := client.FetchNotifications(httpClient)
		if notifErr != nil {
			fmt.Printf("Error fetching notifications: %v\n", notifErr)
		}

		time.Sleep(60 * time.Second)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

}
