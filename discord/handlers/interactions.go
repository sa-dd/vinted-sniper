package handlers

import (
    "log"

    "github.com/bwmarrin/discordgo"
    "vintsnipe/discord/commands"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    switch i.Type {
    case discordgo.InteractionApplicationCommand:
        handleSlashCommand(s, i)
    }
}

func handleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
    commandName := i.ApplicationCommandData().Name
    
    if cmd, exists := commands.CommandRegistry[commandName]; exists {
        cmd.Handler(s, i)
    } else {
        log.Printf("Unknown command: %s", commandName)
        respondWithError(s, i, "Unknown command")
    }
}

func respondWithError(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
		Content: "Error: " + message,
            Flags:   discordgo.MessageFlagsEphemeral,
        },
    })
}

