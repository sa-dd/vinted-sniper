package commands

import (
	"github.com/bwmarrin/discordgo"
	"vintsnipe/discord/types"
)

var CommandRegistry = map[string]types.Command{
	"search": {
		Definition: Search,
		Handler:    SearchHandler,
	},
	"filter": {
		Definition: Filter,
		Handler:    FilterHandler,
	},
}

func GetAllCommands() []*discordgo.ApplicationCommand {
	var commands []*discordgo.ApplicationCommand
	for _, cmd := range CommandRegistry {
		commands = append(commands, cmd.Definition)
	}
	return commands
}
