package telegram

import tb "gopkg.in/tucnak/telebot.v2"

type Message struct {
	Original *tb.Message
}

func NewMessage(message *tb.Message) *Message {
	return &Message{
		Original: message,
	}
}
