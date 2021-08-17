package discord

import (
	"strconv"

	"github.com/diamondburned/arikawa/discord"
	"github.com/leviska/batya-go/batya"
)

type ID discord.Snowflake

func (ID) Source() string {
	return NetworkName
}

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (n *Network) ToPhotoAdapter(m *discord.Attachment) *batya.Photo {
	return &batya.Photo{
		URL:      m.URL,
		Width:    int(m.Width),
		Height:   int(m.Height),
		Original: m,
	}
}

func (n *Network) ToMessageAdapter(m *discord.Message) *batya.Message {
	message := &batya.Message{
		Text: &batya.Text{Text: m.Content},
		Sender: &batya.User{
			ID:   ID(m.Author.ID),
			Name: m.Author.Username,
		},
		SourceID:   ID(m.ChannelID),
		SourceName: NetworkName,
		Original:   m,
	}
	if len(m.Attachments) > 0 {
		photo := n.ToPhotoAdapter(&m.Attachments[0])
		if photo != nil {
			message.Photo = photo
		}
	}
	return message
}

func (n *Network) FromPhotoAdapter(p *batya.Photo) *discord.Embed {
	embed := discord.NewEmbed()
	embed.Image = &discord.EmbedImage{
		URL: p.URL,
	}
	return embed
}

func (n *Network) FromMessageAdapter(m *batya.Message) (string, *discord.Embed) {
	var text string
	if m.Text != nil {
		text = m.Text.Text
	}
	var embed *discord.Embed
	if m.Photo != nil {
		embed = n.FromPhotoAdapter(m.Photo)
	}
	return text, embed
}
