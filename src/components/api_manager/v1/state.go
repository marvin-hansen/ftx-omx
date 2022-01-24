// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"context"
	"sync"
)

type State struct {
	sync.RWMutex
	cMap map[string]context.CancelFunc
}

func newState() (state *State) {
	state = &State{
		cMap: make(map[string]context.CancelFunc),
	}
	return state
}
