// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/clients/pgdb"
	"ftx-omx/src/components/account_manager"
	"ftx-omx/src/components/api_manager"
	"ftx-omx/src/components/component_manager"
	"ftx-omx/src/components/order_manager"
	"ftx-omx/src/types"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
	"time"
)

type State struct {
	config        *types.MainConfig
	dbConf        *pgdb.DBConfig
	dbModel       []interface{}
	compManager   component_manager.ComponentManager
	apiManager    api_manager.ApiManager
	clientManager account_manager.AccountManager
	orderManager  order_manager.OrderManager
	dbComp        *pgdb.DBComponent
	serviceID     string
	serviceName   string
	port          string
	network       string
	httpServer    *http.Server
	lis           net.Listener
	httpListener  net.Listener
	mux           cmux.CMux
	startTime     time.Time
	ready         bool
	apiOn         bool
	orderOn       bool
}

func NewState(mainConfig *types.MainConfig) (state *State) {
	state = &State{
		config:        mainConfig,
		dbConf:        mainConfig.DBConf,
		dbModel:       mainConfig.DBConf.DBModel,
		serviceID:     mainConfig.ServiceID,
		serviceName:   mainConfig.ServiceName,
		port:          mainConfig.Port,
		network:       mainConfig.Network,
		compManager:   nil,
		dbComp:        nil,
		apiManager:    nil,
		clientManager: nil,
		orderManager:  nil,
		httpServer:    nil,
		lis:           nil,
		httpListener:  nil,
		mux:           nil,
		startTime:     time.Now(),
		ready:         false,
		apiOn:         mainConfig.ApiOn,
		orderOn:       mainConfig.OrderOn,
	}
	return state
}
