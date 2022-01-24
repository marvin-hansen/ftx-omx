// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/clients/pgdb"
	"ftx-omx/src/components/crypto_manager"
)

type DataManager struct {
	dependencies *Dependencies
	state        *State
}

func NewDataManager(cryptoManager crypto_manager.CryptoManager, dbComp *pgdb.DBComponent) *DataManager {
	// CIRA = Construction = Initialization = Return (Resource) Allocation
	comp := DataManager{
		newDependencies(cryptoManager, dbComp),
		newState()} // 1. Construction
	comp.init()  // 2. Initialization
	return &comp // 3. Return Reference to (Resource) Allocation
}
