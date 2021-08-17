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

func (ID) Source() string {
	return "telegram"
}

func (id ID) Recipient() string {
	return id.String()
}

func (n *Network) ToPhotoAdapter(p *tb.Photo) *batya.Photo {
	url, err := n.Bot.FileURLByID(p.FileID)
	if err != nil {
		return nil
	}
	return &batya.Photo{
		URL:      url,
		Width:    p.Width,
		Height:   p.Height,
		Original: p,
	}
}

func (n *Network) ToMessageAdapter(m *tb.Message) *batya.Message {
	message := &batya.Message{
		Text: &batya.Text{Text: m.Text},
		Sender: &batya.User{
			ID:   ID(m.Sender.ID),
			Name: m.Sender.FirstName,
		},
		SourceID:   ID(m.Chat.ID),
		SourceName: NetworkName,
		Original:   m,
	}
	if m.Photo != nil {
		photo := n.ToPhotoAdapter(m.Photo)
		if photo != nil {
			message.Photo = photo
		}
	}
	return message
}

type Text string

func (t Text) Send(bot *tb.Bot, to tb.Recipient, options *tb.SendOptions) (*tb.Message, error) {
	return bot.Send(to, string(t))
}

func (n *Network) FromMessageAdapter(m *batya.Message) tb.Sendable {
	if m.Photo != nil {
		var photo *tb.Photo
		if tgPhoto, ok := m.Photo.Original.(*tb.Photo); ok {
			photo = tgPhoto
		} else {
			photo = &tb.Photo{
				File: tb.FromURL(m.Photo.URL),
			}
		}
		if m.Text != nil {
			photo.Caption = m.Text.Text
		}
		return photo
	} else {
		return Text(m.Text.Text)
	}
}
