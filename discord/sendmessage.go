package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendNotif(dg *discordgo.Session, embeds []*discordgo.MessageEmbed) {

	s.ChannelMessageSend("1415381963367383165", "NiggaThrum")
	log.Println(len(embeds))
	if len(embeds) < 10 {
		_, err := s.ChannelMessageSendEmbeds("1415381963367383165", embeds)
		if err != nil{
			log.Println("Error sending Embeds: ", err)
		}
	}
	//log.Println(mg)
}
