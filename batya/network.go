package batya

type TextCallback func(Network, *Message)

type Network interface {
	Sourcer
	HandleText(callback TextCallback)
	SendMessage(to ID, message *Message) error
	Start() error
}
