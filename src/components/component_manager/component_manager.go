// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package component_manager

import (
	"ftx-omx/src/clients/pgdb"
	"ftx-omx/src/components/account_manager"
	"ftx-omx/src/components/api_manager"
	v1 "ftx-omx/src/components/component_manager/v1"
	"ftx-omx/src/components/crypto_manager"
	"ftx-omx/src/components/data_manager"
	"ftx-omx/src/components/order_manager"
)

type ComponentManager interface {
	GetAccountManager() account_manager.AccountManager
	GetApiManager() api_manager.ApiManager
	GetCryptoManager() crypto_manager.CryptoManager
	GetDataManager() data_manager.DataManager
	GetOrderManager() order_manager.OrderManager
}

func NewComponentManager(dbComp *pgdb.DBComponent) ComponentManager {
	return v1.NewComponentManager(dbComp)
}
