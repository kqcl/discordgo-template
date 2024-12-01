package util

import "github.com/bwmarrin/discordgo"

var Color = struct{
    Success int
    Error int
    Warning int
    Info int
}{
 	Success: 0x57F287, // discord green
	Error:   0xED4245, // discord red
	Warning: 0xFEE75C, // discord yellow
	Info:    0x5865F2, // discord blue
}

func SimpleTextResponse(content string, s *discordgo.Session, i *discordgo.InteractionCreate) error {
    err := s.InteractionRespond(
        i.Interaction,
        &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data : &discordgo.InteractionResponseData {
                Content: content,
            },
        },
    )
    return err
}

func SimmpleResponse(content string, s *discordgo.Session, i *discordgo.InteractionCreate) error {
    embed := &discordgo.MessageEmbed{
        Color: Color.Success,
        Description: content,
    }
    err := s.InteractionRespond(
        i.Interaction,
        &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData {
                Embeds: []*discordgo.MessageEmbed{embed},
            },
        },
    )
    return err
}


func SimpleError(content string, s *discordgo.Session, i *discordgo.InteractionCreate) error {
    embed := &discordgo.MessageEmbed{
        Color: Color.Error,
        Description: content,
    }
    err := s.InteractionRespond(
        i.Interaction,
        &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData {
                Embeds: []*discordgo.MessageEmbed{embed},
            },
        },
    )
    return err
}
