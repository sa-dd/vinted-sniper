package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"vintsnipe"
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
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "filter",
			Description: "Add relevent filters to item categories",
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "price",
					Value: "price",
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionNumber,
			Name:        "amount",
			Description: "Any item greater than this amount will be omitted",
			Required:    true,
		},
	},
}

func SearchHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	var channelID string

	options := i.ApplicationCommandData().Options

	guildId := s.State.Ready.Guilds[0].ID
	createChannel := options[1].Value.(bool)
	search := strings.Replace(options[0].Value.(string), " ", "%20", 1)
	channelName := options[0].Value.(string)

	if createChannel {
		channelName = options[0].Value.(string) + " pol"
		ch, _ := s.GuildChannelCreateComplex(guildId, discordgo.GuildChannelCreateData{
			Name:     channelName,
			Type:     discordgo.ChannelTypeGuildText,
			ParentID: "1423023839847256155",
		})

		channelID = ch.ID
	}

	filter := &vintsnipe.Filter{}
	filter.Type = options[2].Value.(string)

	switch filter.Type {
	case "price":
		filter.Value = options[3].FloatValue()

	default:
		filter = nil
	}

	pol, err := vintsnipe.Create(filter, channelID, search)
	if err != nil {
		fmt.Errorf("failed to create a new poll: ", err)
		return
	}

	vintsnipe.PolSlice = append(vintsnipe.PolSlice, pol)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("New poll added to %s channel with %s filter set to %v", channelName, filter.Type, filter.Value.(float64)),
		},
	})

}
