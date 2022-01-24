// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/utils/dbg"
)

func (c *OrderManager) init() {
	c.verifyInit()
}

func (c *OrderManager) verifyInit() {
	dbg.NilCheck(c.dependencies.accountManager, "NPE: accountManager is nil. Fix init.")
	dbg.NilCheck(c.dependencies.dataManager, "NPE: dataManager is nil. Fix init.")
	dbg.NilCheck(c.state.orderMap, "NPE: orderMap is nil. Fix init.")
	dbg.NilCheck(c.state.bookOrderMap, "NPE: bookOrderMap is nil. Fix init.")
}
