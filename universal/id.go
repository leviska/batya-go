package universal

import "github.com/leviska/batya-go/batya"

type IDMap map[string]batya.ID

type ID struct {
	idMap IDMap
}

func (id ID) String() string {
	return "unimplement"
}

func (id ID) Source() string {
	return NetworkName
}
