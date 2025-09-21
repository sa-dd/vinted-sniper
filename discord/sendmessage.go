package discord
<<<<<<< HEAD
=======

import (

	"log"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func SendNotif(embeds []*discordgo.MessageEmbed){
	
	dg, err := discordgo.New("Token here")
	if err != nil {
		log.Println("Error initializing Bot instance: ", err)
	}

	dg.AddHandler(func (s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == "shazam" {
			s.ChannelMessageSend("1415381963367383165", "NiggaThrum")
			log.Println(embeds)
			mg, err := s.ChannelMessageSendEmbeds("1415381963367383165", embeds[0:9])
			if err != nil {
				log.Println("Error sending Embeds: ", err)
			}
			log.Println(mg)

		}
	})

	err = dg.Open()
	if err != nil {
		log.Println("Error opening websocket to Gateway API: ", err)
	}
	
	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
>>>>>>> 7c0cb93 (chore: add secrets placeholders)
