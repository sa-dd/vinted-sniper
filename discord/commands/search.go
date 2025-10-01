package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var Search = &discordgo.ApplicationCommand{
	Name:        "search",
	Description: "Search for a specific category/item",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "item",
			Description: "Item to search for",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionBoolean,
			Name:        "new_channel",
			Description: "Whether to update the current poll or create a new one",
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "yes",
					Value: true,
				},
				{
					Name:  "no",
					Value: false,
				},
			},
		},
	},
}

func SearchHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options

	guildId := s.State.Ready.Guilds[0].ID
	channelName := options[0].Value.(string) + " pol"
	flag := options[1].Value.(bool)
	fmt.Println(guildId)
	if flag {
		s.GuildChannelCreateComplex(guildId, discordgo.GuildChannelCreateData{
			Name:     channelName,
			Type:     discordgo.ChannelTypeGuildText,
			ParentID: "1423023839847256155",
		})
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "New search url created",
		},
	})

}
