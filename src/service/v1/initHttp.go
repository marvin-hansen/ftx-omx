// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/utils/dbg"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
)

func (s *Service) initNetwork(msg string) {
	PrintInitHeader(s.state.serviceID, msg)
	s.initHttpNetwork()
	DbgPrint(" * Http network configured.")

	s.initHttpServer()
	DbgPrint(" * Http server configured.")

	s.initHttpRestApi()
	DbgPrint(" * Http REST API configured.")
}

func (s *Service) initHttpNetwork() {
	DbgPrint(" * Init: Network")
	var err error
	s.state.lis, err = net.Listen(s.state.network, s.state.port)
	dbg.LogFatal(err)
	DbgPrint(" * net listener configured")

	s.state.mux = cmux.New(s.state.lis)
	DbgPrint(" * cmux configured")

	s.state.httpListener = s.state.mux.Match(cmux.Any()) // All the rest is assumed to be HTTP
	DbgPrint(" * http listener configured")
}

func (s *Service) initHttpServer() {
	DbgPrint(" * Init: Http Service!")
	// replace with fasthttp
	// https://github.com/valyala/fasthttp#http-server-performance-comparison-with-nethttp
	// https://github.com/valyala/fasthttp/blob/master/examples/helloworldserver/helloworldserver.go
	s.state.httpServer = &http.Server{}
	DbgPrint(" * http server configured")

}

func (s *Service) initHttpRestApi() {
	s.initSysApi()
	DbgPrint(" * System status REST API configured")

	s.initAccountApi()
	DbgPrint("  API Management REST API configured")

	s.initOrderApi()
	DbgPrint(" * Order Routing REST API configured")
}
