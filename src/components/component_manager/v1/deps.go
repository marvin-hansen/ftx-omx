// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/clients/pgdb"
	"web_socket/src/components/account_manager"
	"web_socket/src/components/api_manager"
	"web_socket/src/components/crypto_manager"
	"web_socket/src/components/data_manager"
	"web_socket/src/components/order_manager"
)

type Dependencies struct {
	apiManager     api_manager.ApiManager
	accountManager account_manager.AccountManager
	cryptoManager  crypto_manager.CryptoManager
	dataManager    data_manager.DataManager
	dbComp         *pgdb.DBComponent
	orderManager   order_manager.OrderManager
}

func newDependencies(dbComp *pgdb.DBComponent) (deps *Dependencies) {
	deps = &Dependencies{
		apiManager:     nil,
		accountManager: nil,
		cryptoManager:  nil,
		dataManager:    nil,
		dbComp:         dbComp,
		orderManager:   nil,
	}
	return deps
}
