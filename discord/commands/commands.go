package commands

import (
    "vinted-sniper/discord/types"
    "github.com/bwmarrin/discordgo"

)

var CommandRegistry = map[string]types.Command{
	"search":{
		Definition: Search,
		Handler: SearchHandler,
	},
}

func GetAllCommands() []*discordgo.ApplicationCommand {
    var commands []*discordgo.ApplicationCommand
    for _, cmd := range CommandRegistry {
        commands = append(commands, cmd.Definition)
    }
    return commands
}
