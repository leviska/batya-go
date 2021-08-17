package batya

type Message struct {
	// Data
	Text  *Text
	Photo *Photo
	// Meta
	Sender   *User
	SourceID ID
	// Origin
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
