// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/utils/dbg"
)

func (c *ApiManager) init() {
	c.verifyInit()
}

func (c *ApiManager) verifyInit() {
	dbg.NilCheck(c.state.cMap, "NPE: state.cMap is nil. Fix init.")
}
