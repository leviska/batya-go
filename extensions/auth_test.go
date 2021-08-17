package extensions_test

import (
	"log"
	"testing"

	"github.com/leviska/batya-go/batya"
	"github.com/leviska/batya-go/extensions"
	"github.com/leviska/batya-go/test"
	"github.com/leviska/batya-go/universal"
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
			to := universal.NewID()
			for k, id := range uni.IDs {
				if k != m.Source() {
					to.IDs[k] = id
				}
			}
			err = n.SendMessage(to, m)
		}
		if err != nil {
			log.Panic(err)
		}
	})
	network.Start()
}
