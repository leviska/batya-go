package telegram

import (
	"strconv"

	"github.com/leviska/batya-go/batya"
	tb "gopkg.in/tucnak/telebot.v2"
)

type ID int

func (id ID) String() string {
	return strconv.Itoa(int(id))
}

func (id ID) Recipient() string {
	return id.String()
}

func MessageAdapter(m *tb.Message) *batya.Message {
	return &batya.Message{
		Text: &batya.Text{Text: m.Text},
		Sender: &batya.User{
			ID: ID(m.Sender.ID),
			Name: m.Sender.FirstName,
		},
		Original: m,
	}
}
