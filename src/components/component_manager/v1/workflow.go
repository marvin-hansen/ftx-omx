// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/components/account_manager"
	"web_socket/src/components/api_manager"
	"web_socket/src/components/crypto_manager"
	"web_socket/src/components/data_manager"
	"web_socket/src/components/order_manager"
)

func (c *ComponentManager) GetApiManager() api_manager.ApiManager {
	return c.dependencies.apiManager
}

func (c *ComponentManager) GetAccountManager() account_manager.AccountManager {
	return c.dependencies.accountManager
}

func (c *ComponentManager) GetCryptoManager() crypto_manager.CryptoManager {
	return c.dependencies.cryptoManager
}

func (c *ComponentManager) GetDataManager() data_manager.DataManager {
	return c.dependencies.dataManager
}

func (c *ComponentManager) GetOrderManager() order_manager.OrderManager {
	return c.dependencies.orderManager
}
