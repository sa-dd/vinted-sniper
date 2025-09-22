package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendNotif(dg *discordgo.Session, embeds []*discordgo.MessageEmbed) {

	dg.ChannelMessageSend("1415381963367383165", "NiggaThrum")
	//log.Println("Sending : ",  len(embeds))
	if len(embeds) < 10 {
		_, err := dg.ChannelMessageSendEmbeds("1415381963367383165", embeds)
		if err != nil{
			log.Println("Error sending Embeds: ", err)
		}
	}
	//log.Println(mg)
}
