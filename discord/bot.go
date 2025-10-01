package discord

import (
	"log"
	"github.com/bwmarrin/discordgo"
)

func Init() (*discordgo.Session, error) {

	dg, err := discordgo.New("DISCORD_TOKEN")
	if err != nil {
		log.Println("Error initializing Bot instance: ", err)
	}

	err = dg.Open()
	if err != nil {
		log.Println("Error opening websocket to Gateway API: ", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")


	return dg, err

}
