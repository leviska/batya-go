package extensions

import (
	"unicode"

	"github.com/leviska/batya-go/batya"
)

type routerMap map[string]batya.TextCallback

type Router struct {
	handlers    routerMap
	textHandler batya.TextCallback
}

func NewRouter(network batya.Network, textHandler batya.TextCallback) *Router {
	router := &Router{
		handlers:    routerMap{},
		textHandler: textHandler,
	}
	network.HandleText(func(n batya.Network, m *batya.Message) {
		router.handleText(n, m)
	})
	return router
}

func (r *Router) Handle(command string, handler batya.TextCallback) {
	r.handlers[command] = handler
}

func (r *Router) Unhandle(command string) {
	delete(r.handlers, command)
}

func findSpace(str string) int {
	for pos, r := range str {
		if unicode.IsSpace(r) {
			return pos
		}
	}
	return len(str)
}

func (r *Router) tryCommand(n batya.Network, m *batya.Message) bool {
	text := m.Text.Text
	if len(text) < 2 {
		return false
	}
	if text[0] != '/' {
		return false
	}
	text = text[1:]

	pos := findSpace(text)
	command := text[:pos]
	handler := r.handlers[command]
	if handler == nil {
		return false
	}
	m.Text.Text = text[pos:]
	
	handler(n, m)
	return true
}

func (r *Router) handleText(n batya.Network, m *batya.Message) {
	if !r.tryCommand(n, m) {
		if r.textHandler != nil {
			r.textHandler(n, m)
		}
	}
}
