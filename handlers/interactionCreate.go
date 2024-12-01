package handlers

import (
	"handler/cmd"
	"log"

	"github.com/bwmarrin/discordgo"
)

func SlashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    data := i.ApplicationCommandData()

    for _, command := range cmd.CommandRegistry {
        if data.Name == command.Name {
            if err := command.Callback(s, i); err != nil {
                log.Printf("error executing command %s: %v", data.Name, err)
            }
            return
        }
    }
    log.Printf("command %s not found", data.Name)
}
