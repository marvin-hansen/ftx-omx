// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package api_manager

import (
	v1 "ftx-omx/src/components/api_manager/v1"
	"ftx-omx/src/components/data_manager"
	"ftx-omx/src/types"
)

type ApiManager interface {
	AddApi(value types.Api)
	GetApi(key string) (types.Api, error)
	UpdateApi(key string, newValue types.Api)
	CheckApiExists(key string) (ok bool)
	DeleteApi(key string)
	StartMonitorApi(api types.Api)
	StopMonitorApi(key string)
	StartAllMonitorApis()
	StopAllMonitorApis()
}

func NewApiManager(dataManager data_manager.DataManager) (apiManager ApiManager) {
	apiManager = v1.NewApiManager(dataManager)
	return apiManager
}
