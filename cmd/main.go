package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vintsnipe/client"
	"vintsnipe/discord"
	"vintsnipe/discord/handlers"
	"vintsnipe/global"
)

func main() {

	dg, err := discord.Init()
	if err != nil {
		log.Println("Error initializing Bot instance: ", err)
	}

	for {
		//fmt.Println(global.PolSlice)

		for _, i := range global.PolSlice {
			items, err := client.FetchItems(i.Client, i.Request)
			if err != nil {
				fmt.Printf("Error fetching items: %v\n", err)
			}

			if i.LatestItemID != 0 {
				latestItems := client.FindLatestItems(i.LatestItemID, items)
				fmt.Println(fmt.Sprintf("Latest item ID: %d", i.LatestItemID))
				fmt.Println(fmt.Sprintf("Latest Items: %d", len(latestItems)))
				if len(latestItems) != 0 {
					notifs := discord.CreateEmbed(latestItems)
					discord.SendNotif(dg, notifs, i.ChannelID)
				}
			}

			if len(items) > 0 {
				i.LatestItemID = items[0].Id
			}

		time.Sleep(60 * time.Second)
		}

	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	handlers.UnregisterSlashCommands(dg, "")

}
