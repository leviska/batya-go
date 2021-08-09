package telegram

import (
	"log"
	"os"
	"testing"

	"github.com/leviska/batya-go/batya"
	tb "gopkg.in/tucnak/telebot.v2"
)

func TestPingPong(t *testing.T) {
	token := os.Getenv("TELEGRAM_TOKEN")
	network, err := NewTelegram(token)
	if err != nil {
		log.Fatal(err)
		return
	}

	network.HandleText(func(n batya.Network, m batya.Message) {
		mTG := m.(*tb.Message)
		n.SendMessage(mTG.Sender, mTG.Text)
	})

	network.Start()
}
