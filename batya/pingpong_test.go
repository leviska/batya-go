package batya_test

import (
	"log"
	"os"
	"sync"
	"testing"

	"github.com/leviska/batya-go/batya"
	"github.com/leviska/batya-go/discord"
	"github.com/leviska/batya-go/telegram"
)

func TestPingPong(t *testing.T) {
	token := os.Getenv("TELEGRAM_TOKEN")
	tgNetwork, err := telegram.NewTelegram(token)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	token = os.Getenv("DISCORD_TOKEN")
	dsNetwork, err := discord.NewDiscord(token)
	if err != nil {
		log.Fatal(err)
		return
	}

	networks := []batya.Network{ tgNetwork, dsNetwork }
	
	wait := sync.WaitGroup{}
	for _, ntw := range networks {
		ntw := ntw

		ntw.HandleText(func(n batya.Network, m *batya.Message) {
			n.SendMessage(m.SourceID, m)
		})
		
		wait.Add(1)
		go func() {
			defer wait.Done()
			err := ntw.Start()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wait.Wait()
}
