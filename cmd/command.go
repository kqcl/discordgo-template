package cmd

import "github.com/bwmarrin/discordgo"

type Command struct {
    Name string
    Description string
    Callback func(s *discordgo.Session, i *discordgo.InteractionCreate) error
    Options []discordgo.ApplicationCommandOption
}

var CommandRegistry []Command

func RegisterCommand(command Command) {
    CommandRegistry = append(CommandRegistry, command)
}
