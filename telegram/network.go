package telegram

import (
	"time"

	"github.com/leviska/batya-go/batya"
	tb "gopkg.in/tucnak/telebot.v2"
)

const NetworkName = "telegram"

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

func (n *Network) Source() string {
	return NetworkName
}

func (n *Network) HandleText(callback batya.TextCallback) {
	n.Bot.Handle(tb.OnText, func(message *tb.Message) {
		callback(n, MessageAdapter(message))
	})
}

func (n *Network) SendMessage(to batya.ID, message *batya.Message) error {
	_, err := n.Bot.Send(to.(ID), message.Text.Text)
	return err
}

func (n *Network) Start() error {
	n.Bot.Start()
	return nil
}
