package telegram

import (
	"log"
	"os"
	"testing"

	"github.com/leviska/batya-go/batya"
)

func TestPingPong(t *testing.T) {
	token := os.Getenv("TELEGRAM_TOKEN")
	network, err := NewTelegram(token)
	if err != nil {
		log.Fatal(err)
		return
	}

	network.HandleText(func(n batya.Network, m batya.Message) {
		mTG := m.(*Message)
		n.SendMessage(mTG.Original.Sender, mTG.Original.Text)
	})

	network.Start()
}