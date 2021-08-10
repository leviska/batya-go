package discord

import (
	"github.com/diamondburned/arikawa/gateway"
	"github.com/leviska/batya-go/batya"
)

func MessageAdapter(m *gateway.MessageCreateEvent) *batya.Message {
	return &batya.Message{
		Text: &batya.Text{Text: m.Content},
		Sender: &batya.User{
			ID: m.Author.ID, 
			Name: m.Author.Username,
		},
		SourceID: m.ChannelID,
		Source: "discord",
		Original: m,
	}
}
