package batya

type Message struct {
	Text     *Text
	Photos   []*Photo
	Sender   *User
	SourceID ID

	SourceName string
	Original   interface{}
}

func NewTextMessage(text string) *Message {
	return &Message{
		Text: &Text{
			Text: text,
		},
	}
}

func (m *Message) Source() string {
	return m.SourceName
}
