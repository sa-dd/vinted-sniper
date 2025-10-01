package commands

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
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
    },
    
}

func SearchHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options
	fmt.Println(options)

	// make a new url with the given params

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "New search url created",
        },
    })

}
