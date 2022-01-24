// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"github.com/go-numb/go-ftx/rest"
	"sync"
	"web_socket/src/types"
)

type State struct {
	sync.RWMutex
	clientMap *types.SyncedOrderedMap[string, *rest.Client]
}

func newState() (state *State) {
	state = &State{
		clientMap: types.NewSyncedOrderedMap[string, *rest.Client](),
	}
	return state
}
