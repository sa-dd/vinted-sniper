package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var Filter = &discordgo.ApplicationCommand{
	Name:        "filter",
	Description: "Add Filters for a specific item",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
			Name:        "price",
			Description: "Set a filter on the Price of the Item",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "less",
					Description: "Item price should be less than the specified amount",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionNumber,
							Name:        "amount",
							Description: "Set a filter on the Price of the Item",
							Required:    true,
						},
					},
				},

				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "greater",
					Description: "Item price should be greater than the specified amount",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionNumber,
							Name:        "amount",
							Description: "Set a filter on the Price of the Item",
							Required:    true,
						},
					},
				},
			},
		},
	},
}

func FilterHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options

	//guildId := s.State.Ready.Guilds[0].ID
	fmt.Println(options[0].Options[0].Options[0].FloatValue())

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Filter Set",
		},
	})

}
