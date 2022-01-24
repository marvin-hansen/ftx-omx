// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

func (c *CryptoManager) init() {
	c.verifyInit()
}

func (c *CryptoManager) verifyInit() {
	//dbg.NilCheck(c.dependencies, "NPE: ... is nil. Fix init.")
}
