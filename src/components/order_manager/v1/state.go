// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/types"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"sync"
)

type State struct {
	sync.RWMutex
	orderMap     *types.SyncedOrderedMap[string, *orders.ResponseForPlaceOrder]
	bookOrderMap *types.SyncedOrderedMap[string, *orders.ResponseForPlaceOrder]
}

func newState() (state *State) {
	state = &State{
		orderMap:     types.NewSyncedOrderedMap[string, *orders.ResponseForPlaceOrder](),
		bookOrderMap: types.NewSyncedOrderedMap[string, *orders.ResponseForPlaceOrder](),
	}
	return state
}
