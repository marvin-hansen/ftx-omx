// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

type CryptoManager struct{}

func NewCryptoManager() *CryptoManager {
	// CIRA = Construction = Initialization = Return (Resource) Allocation
	comp := CryptoManager{} // 1. Construction
	comp.init()             // 2. Initialization
	return &comp            // 3. Return Reference to (Resource) Allocation
}
