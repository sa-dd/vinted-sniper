package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendNotif(dg *discordgo.Session, embeds []*discordgo.MessageEmbed, channelID string) {

	if len(embeds) < 10 {
		_, err := dg.ChannelMessageSendEmbeds(channelID, embeds)
		if err != nil {
			log.Println("Error sending Embeds: ", err)
		}
	}
}
