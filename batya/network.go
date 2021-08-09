package batya

type TextCallback func(Network, Message)

type Network interface {
	HandleText(callback TextCallback)
	SendMessage(to User, message Message) error
}
