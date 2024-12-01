package main

import (
	"handler/handlers"
	"handler/helpers"
	"handler/util"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
    config, err := util.GetConfig()
    if err != nil {
        log.Fatalf("error loading config: %v", err)
    }

    dg, _ := discordgo.New("Bot " + config.Token)

    if err := helpers.RegisterCommands(dg, config.ApplicationID, config.GuildID); err != nil {
        log.Fatalf("error registering commands: %v", err)
    }

    dg.AddHandler(handlers.SlashCommandHandler)

    if err := dg.Open(); err != nil {
        log.Fatalf("error opening session: %v", err)
    }

    log.Printf("Logged in as %s", dg.State.User.Username)

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    log.Println("Bot is listening...")
    <- stop

    if err := dg.Close(); err != nil {
        log.Fatalf("error shutting down gracefully: %v", err)
    }
}
