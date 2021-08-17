package batya

type Message struct {
	Text     *Text
	Photos   []*Photo
	Sender   *User
	SourceID ID

	SourceName string
	Original   interface{}
}

func (m *Message) Source() string {
	return m.SourceName
}
