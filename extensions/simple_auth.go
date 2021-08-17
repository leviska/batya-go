package extensions

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/leviska/batya-go/batya"
	"github.com/leviska/batya-go/universal"
)

type authIndex map[string]map[string]*universal.ID
type authCodeIndex map[string]batya.ID

const authCommand = "auth"

type SimpleAuther struct {
	index     authIndex
	codeIndex authCodeIndex
	router    *Router
}

func NewSimpleAuth(router *Router) *SimpleAuther {
	return &SimpleAuther{
		index:     authIndex{},
		codeIndex: authCodeIndex{},
		router:    router,
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (a *SimpleAuther) Start() {
	a.router.Handle(authCommand, func(n batya.Network, m *batya.Message) {
		var err error
		if m.Text.Text == "" {
			key := RandString(4)
			a.codeIndex[key] = m.SourceID
			err = n.SendMessage(m.SourceID, batya.NewTextMessage(fmt.Sprintf("Send next line to bot where you want to authenticate:\n/auth %s", key)))
		} else {
			id := a.codeIndex[m.Text.Text]
			if id != nil {
				uni := a.Unite(id, m.SourceID)
				err = n.SendMessage(uni, batya.NewTextMessage("Successfully authorized!"))
			} else {
				err = n.SendMessage(m.SourceID, batya.NewTextMessage("Can't find such key"))
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (a *SimpleAuther) Stop() {
	a.router.Unhandle(authCommand)
}

func (a *SimpleAuther) Get(id batya.ID) *universal.ID {
	ntw, ok := a.index[id.Source()]
	if !ok {
		return nil
	}
	uni, ok := ntw[id.String()]
	if !ok {
		return nil
	}
	return uni
}

func (a *SimpleAuther) set(id batya.ID, uni *universal.ID) {
	uni.IDs[id.Source()] = id
	ntw, ok := a.index[id.Source()]
	if !ok {
		ntw = map[string]*universal.ID{}
		a.index[id.Source()] = ntw
	}
	ntw[id.String()] = uni
}

func (a *SimpleAuther) createUser(id batya.ID) *universal.ID {
	return &universal.ID{IDs: universal.IDMap{}}
}

func (a *SimpleAuther) Create(id batya.ID) *universal.ID {
	uni := a.createUser(id)
	a.set(id, uni)
	return uni
}

func (a *SimpleAuther) delete(id batya.ID) {
	ntw, ok := a.index[id.Source()]
	if !ok {
		return
	}
	delete(ntw, id.String())
}

func (a *SimpleAuther) Unite(x batya.ID, y batya.ID) *universal.ID {
	xUni := a.Get(x)
	yUni := a.Get(y)

	if yUni == nil {
		if xUni == nil {
			xUni = a.Create(x)
		}
		a.set(y, xUni)
		return xUni
	} else {
		if xUni == nil {
			a.set(x, yUni)
			return yUni
		} else {
			for _, v := range yUni.IDs {
				a.set(v, xUni)
			}
			return xUni
		}
	}
}

func (a *SimpleAuther) GetOrCreate(id batya.ID) *universal.ID {
	uni := a.Get(id)
	if uni == nil {
		uni = a.Create(id)
	}
	return uni
}
