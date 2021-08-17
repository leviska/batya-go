package extensions

import "github.com/leviska/batya-go/batya"

type authUser map[string]batya.ID
type authUserMap map[int]authUser

type authIndex map[string]map[string]int

type SimpleAuther struct {
	users  authUserMap
	index  authIndex
	lastID int
}

func NewSimpleAuth() *SimpleAuther {
	return &SimpleAuther{
		users: authUserMap{},
		index: authIndex{},
	}
}

