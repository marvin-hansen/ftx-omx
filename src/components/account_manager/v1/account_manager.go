// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import "web_socket/src/components/api_manager"

type AccountManager struct {
	dependencies *Dependencies
	state        *State
}

func NewAccountManager(apiManager api_manager.ApiManager) *AccountManager {
	comp := AccountManager{ // CIRA = Construction = Initialization = Return (Resource) Allocation
		newDependencies(apiManager),
		newState()} // 1. Construction
	comp.init()  // 2. Initialization
	return &comp // 3. Return Reference to (Resource) Allocation
}
