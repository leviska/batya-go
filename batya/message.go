package batya

type Message struct {
	Text     *Text
	Photos   []*Photo
	Sender   *User
	SourceID ID

	Source   string
	Original interface{}
}
