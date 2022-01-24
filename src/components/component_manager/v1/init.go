// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/components/account_manager"
	"ftx-omx/src/components/api_manager"
	"ftx-omx/src/components/crypto_manager"
	"ftx-omx/src/components/data_manager"
	"ftx-omx/src/components/order_manager"
	"ftx-omx/src/utils/dbg"
)

func (c *ComponentManager) init() {
	DbgPrint(" Init crypto_manager! ")
	c.dependencies.cryptoManager = crypto_manager.NewNewCryptoManager()

	DbgPrint(" Init data_manager! ")
	c.dependencies.dataManager = data_manager.NewDataManager(c.dependencies.cryptoManager, c.dependencies.dbComp)

	DbgPrint(" Init api_manager! ")
	c.dependencies.apiManager = api_manager.NewApiManager(c.dependencies.dataManager)

	DbgPrint(" Init account_manager! ")
	c.dependencies.accountManager = account_manager.NewAccountManager(c.dependencies.apiManager)

	DbgPrint(" Init order_manager! ")
	c.dependencies.orderManager = order_manager.NewOrderManager(c.dependencies.accountManager, c.dependencies.dataManager)

	DbgPrint(" Verify Init! ")
	c.verifyInit()
	DbgPrint(" Init Complete! ")
}

func (c *ComponentManager) verifyInit() {
	dbg.NilCheck(c.dependencies.apiManager, "NPE: apiManager is nil. Fix init.")
	dbg.NilCheck(c.dependencies.accountManager, "NPE: accountManager is nil. Fix init.")
	dbg.NilCheck(c.dependencies.cryptoManager, "NPE: cryptoManager is nil. Fix init.")
	dbg.NilCheck(c.dependencies.dbComp, "NPE: dbComp is nil. Fix init.")
	dbg.NilCheck(c.dependencies.dataManager, "NPE: dataManager is nil. Fix init.")
	dbg.NilCheck(c.dependencies.orderManager, "NPE: orderManager is nil. Fix init.")
}
