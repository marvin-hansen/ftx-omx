// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/types"
	"github.com/go-numb/go-ftx/rest"
	"sync"
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
