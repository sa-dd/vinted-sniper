package handlers

import (
    "log"

    "github.com/bwmarrin/discordgo"
    "vintsnipe/discord/commands"
)

func RegisterSlashCommands(s *discordgo.Session, guildID string) {
    allCommands := commands.GetAllCommands()
    
    for _, cmd := range allCommands {
        _, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
        if err != nil {
            log.Printf("Cannot create '%s' command: %v", cmd.Name, err)
        } else {
            log.Printf("Successfully registered command: %s", cmd.Name)
        }
    }
}

func UnregisterSlashCommands(s *discordgo.Session, guildID string) {
    registeredCommands, err := s.ApplicationCommands(s.State.User.ID, guildID)
    if err != nil {
        log.Printf("Could not fetch registered commands: %v", err)
        return
    }

    for _, cmd := range registeredCommands {
        err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
        if err != nil {
            log.Printf("Cannot delete '%s' command: %v", cmd.Name, err)
        } else {
            log.Printf("Successfully deleted command: %s", cmd.Name)
        }
    }
}
