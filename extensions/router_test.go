package extensions_test

import (
	"testing"

	"github.com/leviska/batya-go/batya"
	"github.com/leviska/batya-go/extensions"
	"github.com/leviska/batya-go/test"
)

func TestRouter(t *testing.T) {
	network := test.CreateUniverasl()
	router := extensions.NewRouter(network, func(n batya.Network, m *batya.Message) {
		n.SendMessage(m.SourceID, m)
	})
	router.Handle("hi", func(n batya.Network, m *batya.Message) {
		n.SendMessage(m.SourceID, m)
	})
	network.Start()
}
