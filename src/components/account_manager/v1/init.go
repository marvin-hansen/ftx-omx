// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/utils/dbg"
)

func (c *AccountManager) init() {
	c.verifyInit()
}

func (c *AccountManager) verifyInit() {
	dbg.NilCheck(c.dependencies.apiManager, "NPE: apiManager is nil. Fix init.")
	dbg.NilCheck(c.state.clientMap, "NPE: client map is nil. Fix init.")
}
