package telegram

import (
	"time"

	"github.com/leviska/batya-go/batya"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Network struct {
	Bot *tb.Bot
}

func NewTelegram(token string) (*Network, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: time.Second},
	})
	if err != nil {
		return nil, err
	}
	return &Network{
		Bot: bot,
	}, nil
}

func (n *Network) HandleText(callback batya.TextCallback) {
	n.Bot.Handle(tb.OnText, func(message *tb.Message) {
		callback(n, message)
	})
}

func (n *Network) SendMessage(to batya.User, message batya.Message) error {
	toTG := to.(*tb.User)
	_, err := n.Bot.Send(toTG, message)
	return err
}

func (n *Network) Start() {
	n.Bot.Start()
}
