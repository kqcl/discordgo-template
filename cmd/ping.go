package cmd

import (
	"handler/util"

	"github.com/bwmarrin/discordgo"
)

func init() {
    RegisterCommand(Command{
        Name: "ping",
        Description: "responds with pong",
        Callback: ping,
    })
}

func ping(s *discordgo.Session, i *discordgo.InteractionCreate) error {
    reply := "pong"
    if err := util.SimmpleResponse(reply, s, i); err != nil {
        return err
    }
    return nil
}
