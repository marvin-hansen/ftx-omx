// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/clients/pgdb"
	"web_socket/src/components/crypto_manager"
)

type Dependencies struct {
	cryptoManager crypto_manager.CryptoManager
	dbComp        *pgdb.DBComponent
}

func newDependencies(cryptoManager crypto_manager.CryptoManager, dbComp *pgdb.DBComponent) (deps *Dependencies) {
	deps = &Dependencies{
		cryptoManager,
		dbComp,
	}

	return deps
}
