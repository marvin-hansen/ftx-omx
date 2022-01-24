// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package cfg

const (
	Prod                = true // when true, disables dynamic DB schema re-creation which leads to faster boot time. Set to false during DB development.
	DbgMain             = false
	DbgMemory           = false // read docs/dbg/memory_profile.txt and see src/service/v1/initHttpSys for endpoint
	DbgApiManager       = false
	DbgClientManager    = false
	DbgCryptoManager    = false
	DbgComponentManager = false
	DbgDataManager      = false
	DbgOrderManager     = false
)
