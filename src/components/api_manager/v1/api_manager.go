// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import "ftx-omx/src/components/data_manager"

type ApiManager struct {
	state *State
	deps  *Dependencies
}

func NewApiManager(dataManager data_manager.DataManager) *ApiManager {
	comp := ApiManager{
		newState(),
		newDependencies(dataManager),
	} // 1. Construction
	comp.init()  // 2. Initialization
	return &comp // 3. Return Reference to (Resource) Allocation
}
