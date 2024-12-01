package cmd

import (
	"handler/util"

	"github.com/bwmarrin/discordgo"
)

func init() {
    RegisterCommand(Command{
        Name: "echo",
        Description: "responds with echo",
        Callback: echo,
        Options: []discordgo.ApplicationCommandOption{
            {
                Name: "message",
                Description: "the msg to echo",
                Type: discordgo.ApplicationCommandOptionString,
                Required: false,
            },
        },
    })
}

func echo(s *discordgo.Session, i *discordgo.InteractionCreate) error {
    data := i.ApplicationCommandData()

    var messageOption string

    for _, option := range data.Options {
        switch option.Name {
        case "message":
            messageOption = option.StringValue()
        }
    }

    // example for an error
    if messageOption == "" {
        err := util.SimpleError("please enter a string", s, i)
        return err
    }

    err := util.SimmpleResponse(messageOption, s, i)
    return err
}
