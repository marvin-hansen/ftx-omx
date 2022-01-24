// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package cfg

import (
	"web_socket/src/clients/pgdb"
	"web_socket/src/types"
)

func getDBConfig(e types.Env) (dbConf *pgdb.DBConfig) {
	dbConf = &pgdb.DBConfig{
		Addr:             getDBAddress(e),
		User:             "omxuser",
		Password:         "2ee2e41ec0a7e6441e0038",
		Database:         "omxdb",
		DBModel:          getDBModel(),
		DBCompositeTypes: getDBCompositeTypes(),
	}
	return dbConf
}

const envErr = "Unknown Environment. Select only Dev, Test, Prod"

func getDBAddress(e types.Env) string {
	switch e {
	case types.Dev:
		// ":5432" valid localhost format for postgres
		return ":5432"

	case types.Test:
		// Connect to DB in docker in the same virtual net
		return "timescaledb:5432"

	case types.Prod:
		// Connect to DB in docker in the same virtual net
		return "timescaledb:5432"

	case types.UnknownEnv:
		return envErr

	default:
		return envErr
	}
}

func getDBModel() []interface{} {
	return []interface{}{
		(*types.Api)(nil),
		(*types.Error)(nil),
		(*types.OrderStatus)(nil),
		(*types.OrderFill)(nil),
	}
}

func getDBCompositeTypes() []interface{} {
	return nil
}
