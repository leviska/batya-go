package universal

import "github.com/leviska/batya-go/batya"

type IDMap map[string]batya.ID

type ID struct {
	IDs IDMap
}

func NewID() *ID {
	return &ID{IDs: IDMap{}}
}

func (id *ID) String() string {
	return "unimplement"
}

func (id *ID) Source() string {
	return NetworkName
}
