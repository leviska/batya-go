package batya

type MessageCallback func(Network, *Message)

type Receiver interface {
	Sourcer
	Handle(callback MessageCallback)	
}

type Sender interface {
	Sourcer
	SendMessage(to ID, message *Message) error
}

type SendReceiver interface {
	Sender
	Receiver
}

type Network interface {
	SendReceiver
	Start() error
}
 