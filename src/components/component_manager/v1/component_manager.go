// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import "ftx-omx/src/clients/pgdb"

type ComponentManager struct {
	dependencies *Dependencies
}

func NewComponentManager(dbComp *pgdb.DBComponent) *ComponentManager {
	// CIRA = Construction = Initialization = Return (Resource) Allocation
	comp := ComponentManager{newDependencies(dbComp)} // 1. Construction
	comp.init()                                       // 2. Initialization
	return &comp                                      // 3. Return Reference to (Resource) Allocation
}
