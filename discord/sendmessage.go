package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendNotif(dg *discordgo.Session, embeds []*discordgo.MessageEmbed) {

	dg.ChannelMessageSend("1420459536124215386", "NiggaThrum")
	//log.Println("Sending : ",  len(embeds))
	if len(embeds) < 10 {
		_, err := dg.ChannelMessageSendEmbeds("1420459536124215386", embeds)
		if err != nil{
			log.Println("Error sending Embeds: ", err)
		}
	}
	//log.Println(mg)
}
