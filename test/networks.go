package test

import (
	"log"
	"os"

	"github.com/leviska/batya-go/batya"
	"github.com/leviska/batya-go/discord"
	"github.com/leviska/batya-go/telegram"
	"github.com/leviska/batya-go/universal"
)

func CreateNetworks() []batya.Network {
	token := os.Getenv("TELEGRAM_TOKEN")
	tgNetwork, err := telegram.NewTelegram(token)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	token = os.Getenv("DISCORD_TOKEN")
	dsNetwork, err := discord.NewDiscord(token)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return []batya.Network{tgNetwork, dsNetwork}
}

func CreateUniverasl() *universal.Network {
	return universal.NewNetworks(CreateNetworks())
}
