package batya_test

import (
	"testing"

	"github.com/leviska/batya-go/batya"
	"github.com/leviska/batya-go/test"
)

func TestPingPong(t *testing.T) {
	networks := test.CreateNetworks()

	for _, ntw := range networks {
		ntw.Handle(func(n batya.Network, m *batya.Message) {
			n.SendMessage(m.SourceID, m)
		})
	}

	//test.RunNetworks(networks)
}
