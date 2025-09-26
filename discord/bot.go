package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Initiate() (*discordgo.Session, error) {

	dg, err := discordgo.New("Token Here")
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
