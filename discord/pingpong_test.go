package discord

import (
	"log"
	"os"
	"testing"

	"github.com/leviska/batya-go/batya"
)

func TestPingPong(t *testing.T) {
	token := os.Getenv("DISCORD_TOKEN")
	network, err := NewDiscord(token)
	if err != nil {
		log.Fatal(err)
		return
	}

	network.HandleText(func(n batya.Network, m *batya.Message) {
		n.SendMessage(m.ChatID, m)
	})

	network.Start()
}
