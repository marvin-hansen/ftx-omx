// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import "web_socket/src/clients/pgdb"

type MainConfig struct {
	Environment Env
	DBConf      *pgdb.DBConfig
	ServiceID   string
	ServiceName string
	Port        string
	Network     string
	ApiOn       bool
	OrderOn     bool
}
