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

	network.Handle(func(n batya.Network, m *batya.Message) {
		n.SendMessage(m.SourceID, m)
	})

	network.Start()
}
