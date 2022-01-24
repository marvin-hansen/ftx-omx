// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package cfg

import (
	t "ftx-omx/src/types"
)

// env sets DB host according to selected environment. Dev for localhost, Test for Docker.
const env = t.Prod

func GetMainConfig() *t.MainConfig {
	return &t.MainConfig{
		Environment: env,
		DBConf:      getDBConfig(env),
		ServiceID:   "OMX",
		ServiceName: "Order Management & Execution Service! ",
		Port:        ":80",
		Network:     "tcp",
		ApiOn:       false, // switches the api-API on/off. Disabled by default.
		OrderOn:     true,  // switches the order-routing-API on/off. Enabled by default.
	}
}
