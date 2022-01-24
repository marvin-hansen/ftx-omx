// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package data_manager

import (
	"web_socket/src/clients/pgdb"
	"web_socket/src/components/crypto_manager"
	v1 "web_socket/src/components/data_manager/v1"
	"web_socket/src/types"
)

type DataManager interface {
	StoreApi(api types.Api) (ok bool, msg string)
	GetApi(apiID string) (api types.Api, ok bool, msg string)
	GetAllApis() (api []types.Api, ok bool, msg string)
	DeleteApi(apiID string) (ok bool, msg string)
	StoreError(error *types.Error) (ok bool, msg string)
	StoreOrderStatus(orderStatus *types.OrderStatus) (ok bool, msg string)
	StoreOrderFill(fill *types.OrderFill) (ok bool, msg string)
}

func NewDataManager(cryptoManager crypto_manager.CryptoManager, dbComp *pgdb.DBComponent) (dataManager DataManager) {
	dataManager = v1.NewDataManager(cryptoManager, dbComp)
	return dataManager
}
