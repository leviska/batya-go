package universal

import (
	"fmt"
	"sync"

	"github.com/leviska/batya-go/batya"
)

const NetworkName = "universal"

type networksMap map[string]batya.Network

type Network struct {
	networks networksMap
}

func NewNetworks(networks []batya.Network) *Network {
	n := &Network{
		networks: networksMap{},
	}
	for _, ntw := range networks {
		n.networks[ntw.Source()] = ntw
	}
	return n
}

func (n *Network) Source() string {
	return NetworkName
}

func (n *Network) HandleText(callback batya.TextCallback) {
	for _, ntw := range n.networks {
		ntw.HandleText(callback)
	}
}

func (n *Network) sendMessage(source string, to batya.ID, message *batya.Message) error {
	ntw := n.networks[source]
	if ntw == nil {
		return fmt.Errorf("can't find %q network to send message to", source)
	}
	return ntw.SendMessage(to, message)
}

func (n *Network) SendMessage(to batya.ID, message *batya.Message) error {
	if uniID, ok := to.(ID); ok {
		for source, id := range uniID.idMap {
			return n.sendMessage(source, id, message)
		}
	} else {
		return n.sendMessage(to.Source(), to, message)
	}
	return nil
}

func (n *Network) Start() error {
	wait := sync.WaitGroup{}
	errChan := make(chan error)
	for _, ntw := range n.networks {
		wait.Add(1)
		go func(ntw batya.Network) {
			defer wait.Done()
			err := ntw.Start()
			if err != nil {
				errChan <- err
			}
		}(ntw)
	}
	notifyChan := make(chan struct{})
	go func() {
		wait.Wait()
		notifyChan <- struct{}{}
	}()
	select {
	case <-notifyChan:
		return nil
	case err := <-errChan:
		return err
	}
}

func (n *Network) GetNetwork(name string) batya.Network {
	return n.networks[name]
}
