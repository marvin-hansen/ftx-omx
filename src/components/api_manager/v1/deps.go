// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/components/data_manager"
)

type Dependencies struct {
	dataManager data_manager.DataManager
}

func newDependencies(dataManager data_manager.DataManager) (deps *Dependencies) {
	deps = &Dependencies{
		dataManager: dataManager,
	}
	return deps
}
