// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package component_manager

import (
	"web_socket/src/clients/pgdb"
	"web_socket/src/components/account_manager"
	"web_socket/src/components/api_manager"
	v1 "web_socket/src/components/component_manager/v1"
	"web_socket/src/components/crypto_manager"
	"web_socket/src/components/data_manager"
	"web_socket/src/components/order_manager"
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
