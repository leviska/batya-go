package batya

type Message struct {
	Text   *Text
	Photos []*Photo
	Sender *User
	ChatID ID

	Original interface{}
}
