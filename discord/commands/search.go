package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"vintsnipe"
	"vintsnipe/global"
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

	var channelID string

	options := i.ApplicationCommandData().Options

	guildId := s.State.Ready.Guilds[0].ID
	createChannel := options[1].Value.(bool)
	search := strings.Replace(options[0].Value.(string), " ", "%20", 1)

	if createChannel {
		channelName := options[0].Value.(string) + " pol"
		ch, _ := s.GuildChannelCreateComplex(guildId, discordgo.GuildChannelCreateData{
			Name:     channelName,
			Type:     discordgo.ChannelTypeGuildText,
			ParentID: "1423023839847256155",
		})

		channelID = ch.ID
	}

	pol, _ := vintsnipe.CreatePol(channelID, search)

	global.PolSlice = append(global.PolSlice, pol)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "New search url created",
		},
	})

}
