// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/clients/pgdb"
	"ftx-omx/src/components/account_manager"
	"ftx-omx/src/components/api_manager"
	"ftx-omx/src/components/crypto_manager"
	"ftx-omx/src/components/data_manager"
	"ftx-omx/src/components/order_manager"
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
