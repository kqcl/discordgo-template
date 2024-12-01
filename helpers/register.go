package helpers 

import (
	"handler/cmd"
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(dg *discordgo.Session, applicationID string, guildID string) error {
    var discordCommands []*discordgo.ApplicationCommand

    for _, command := range cmd.CommandRegistry {
        discordCommand := &discordgo.ApplicationCommand{
            Name: command.Name,
            Description: command.Description,
        }
        if len(command.Options) > 0 {
            var discordOptions []*discordgo.ApplicationCommandOption
            for _, option := range command.Options {
                discordOption := &discordgo.ApplicationCommandOption{
                    Name: option.Name,
                    Description: option.Description,
                    Type: option.Type,
                    Required: option.Required,
                }

                discordOptions = append(discordOptions, discordOption)
            }

            discordCommand.Options = discordOptions
        }

        discordCommands = append(discordCommands, discordCommand)
    }


    _, err := dg.ApplicationCommandBulkOverwrite(applicationID, guildID, discordCommands)
    if err != nil {
        return err
    }

    log.Println("cmds successfully registered...")
    return nil
}

