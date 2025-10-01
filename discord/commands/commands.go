package commands

import (
    "vinted-sniper/discord/types"
)

var CommandRegistry map[string]types.Command

func GetAllCommands() []*discordgo.ApplicationCommand {
    var commands []*discordgo.ApplicationCommand
    for _, cmd := range CommandRegistry {
        commands = append(commands, cmd.Definition)
    }
    return commands
}
