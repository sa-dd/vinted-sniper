package discord

import (
	"log"
	"os"
	"github.com/bwmarrin/discordgo"
	"vinted-sniper/discord/handlers"
)

func Init() (*discordgo.Session, error) {
	
	token := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New(token)
	if err != nil {
		log.Println("Error initializing Bot instance: ", err)
	}

	dg.AddHandler(handlers.InteractionHandler)

	err = dg.Open()
	if err != nil {
		log.Println("Error opening websocket to Gateway API: ", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")

	handlers.RegisterSlashCommands(dg, "")


	return dg, err

}
