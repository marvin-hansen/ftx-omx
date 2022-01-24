// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/utils/dbg"
)

func (c *DataManager) init() {
	c.verifyInit()
}

func (c *DataManager) verifyInit() {
	dbg.NilCheck(c.state.apiCache, "NPE: apiCac is nil. Fix init.")
	dbg.NilCheck(c.dependencies.dbComp, "NPE: dbComp is nil. Fix init.")
	dbg.NilCheck(c.dependencies.cryptoManager, "NPE: cryptoManager is nil. Fix init.")
}
