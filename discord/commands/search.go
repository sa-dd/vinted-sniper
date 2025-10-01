package commands

import (
	"github.com/bwmarrin/discordgo"
)

var Search = &discordgo.ApplicationCommand{
	Name: "search",
	Description: "Set search params",
	Options: []*discordgo.ApplicationCommandOption{
		Type: 3,

	}

}

func SearchHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

func SlashInit(dg *discordgo.Session) ([]*discordgo.Application, error) {

	appcmd := discordgo.ApplicationCommand{}
	appcmd.Name = "url"
	appcmd.Description= "Add search params here"
	
	cmd, err := dg.ApplicationCommandCreate(dg.State.User.Id, "", &appcmd)
	if err != nil {

		return nil
	}

	//879789745125855283
}
