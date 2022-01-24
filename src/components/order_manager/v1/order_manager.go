// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/components/account_manager"
	"web_socket/src/components/data_manager"
)

type OrderManager struct {
	dependencies *Dependencies
	state        *State
}

func NewOrderManager(clientManager account_manager.AccountManager, dataManager data_manager.DataManager) *OrderManager {
	// CIRA = Construction = Initialization = Return (Resource) Allocation
	comp := OrderManager{
		newDependencies(clientManager, dataManager),
		newState()} // 1. Construction
	comp.init()  // 2. Initialization
	return &comp // 3. Return Reference to (Resource) Allocation
}
