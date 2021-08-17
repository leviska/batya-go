package discord

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/leviska/batya-go/batya"
)

type ID discord.Snowflake

func (ID) Source() string {
	return NetworkName
}

func (id ID) String() string {
	return id.String() 
}

func MessageAdapter(m *gateway.MessageCreateEvent) *batya.Message {
	return &batya.Message{
		Text: &batya.Text{Text: m.Content},
		Sender: &batya.User{
			ID: ID(m.Author.ID), 
			Name: m.Author.Username,
		},
		SourceID: ID(m.ChannelID),
		SourceName: NetworkName,
		Original: m,
	}
}
