// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/components/account_manager"
	"ftx-omx/src/components/data_manager"
)

type Dependencies struct {
	accountManager account_manager.AccountManager
	dataManager    data_manager.DataManager
}

func newDependencies(clientManager account_manager.AccountManager, dataManager data_manager.DataManager) (deps *Dependencies) {
	deps = &Dependencies{
		accountManager: clientManager,
		dataManager:    dataManager,
	}
	return deps
}
