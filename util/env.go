package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


func getEnv() error {
    if err := godotenv.Load(".env"); err != nil {
        return err
    }
    return nil
}

type BotConfig struct {
    Token string
    ApplicationID string
    GuildID string
}

func GetConfig() (BotConfig, error) {
    if err := getEnv(); err != nil {
        return BotConfig{}, fmt.Errorf("error loading environment (.env): %v", err)
    }

    token, exists := os.LookupEnv("TOKEN")
    if !exists {
        return BotConfig{}, errors.New("token is not set as an environment variable")
    }
    appId, exists := os.LookupEnv("APPLICATION_ID")
    if !exists {
        return BotConfig{}, errors.New("app id is not set as an environment variable")
    }
    guildId, exists := os.LookupEnv("GUILD_ID")
    if !exists {
        return BotConfig{}, errors.New("guild id is not set as an environment variable")
    }

    return BotConfig{
        Token: token,
        ApplicationID: appId,
        GuildID: guildId,
    }, nil
}
