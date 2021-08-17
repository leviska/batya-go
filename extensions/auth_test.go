package extensions_test

import (
	"log"
	"testing"

	"github.com/leviska/batya-go/batya"
	"github.com/leviska/batya-go/extensions"
	"github.com/leviska/batya-go/test"
)


func TestAuth(t *testing.T) {
	network := test.CreateUniverasl()
	router := extensions.NewRouter(network)
	auth := extensions.NewSimpleAuth(router)
	auth.Start()
	router.HandleText(func(n batya.Network, m *batya.Message) {
		uni := auth.Get(m.SourceID)
		var err error
		if uni == nil {
			err = n.SendMessage(m.SourceID, m)
		} else {
			err = n.SendMessage(uni, m)
		}
		if err != nil {
			log.Panic(err)
		}
	})
	network.Start()
}
