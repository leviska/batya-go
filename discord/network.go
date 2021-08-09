package discord

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
	"github.com/leviska/batya-go/batya"
)

type Network struct {
	Bot *session.Session
	Me  *discord.User
}

func NewDiscord(token string) (*Network, error) {
	s, err := session.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	// Add the needed Gateway intents.
	s.Gateway.AddIntent(gateway.IntentGuildMessages)
	s.Gateway.AddIntent(gateway.IntentDirectMessages)
	return &Network{
		Bot: s,
	}, nil
}

func (n *Network) HandleText(callback batya.TextCallback) {
	n.Bot.AddHandler(func(c *gateway.MessageCreateEvent) {
		if c.Author.ID != n.Me.ID {
			callback(n, c)
		}
	})
}

func (n *Network) SendMessage(to batya.User, message batya.Message) error {
	_, err := n.Bot.SendMessage((to.(discord.ChannelID)), message.(string), nil)
	return err
}

func (n *Network) Start() error {
	if err := n.Bot.Open(); err != nil {
		return err
	}
	defer n.Bot.Close()
	me, err := n.Bot.Me()
	if err != nil {
		return err
	}
	n.Me = me
	select {}
}
