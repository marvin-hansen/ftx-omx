// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func (s *Service) Serve() {
	PrintDbgHeader()
	g := errgroup.Group{}

	g.Go(func() error {
		return s.state.httpServer.Serve(s.state.httpListener)
	})
	DbgPrint(" * Http server started")

	g.Go(func() error {
		return s.state.mux.Serve()
	})
	DbgPrint(" * Multiplexer stared!")

	s.state.ready = true // this sets /health handler to return 200 status code
	DbgPrint(" * Service Ready & Running")
	PrintStartHeader(s.state.serviceName, s.state.port, time.Since(s.state.startTime))

	err := g.Wait() // Wait for handlers and check for errors
	if err != nil {
		log.Fatal(err)
	}
}
