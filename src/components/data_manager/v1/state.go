// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"sync"
	"web_socket/src/types"
)

type State struct {
	sync.RWMutex
	apiCache *types.SyncedOrderedMap[string, types.Api]
}

func newState() (state *State) {
	state = &State{
		apiCache: types.NewSyncedOrderedMap[string, types.Api](),
	}
	return state
}
