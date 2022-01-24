// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"net/http"
	"net/http/pprof"
)

func (s *Service) initSysApi() {
	// ********************
	// System Status Endpoint
	// ********************

	// curl http://localhost/
	http.HandleFunc("/", s.rootHandler)

	// curl http://localhost/start
	http.HandleFunc("/start", s.startHandler)

	// curl http://localhost/status
	http.HandleFunc("/status", s.statusHandler)

	// curl http://localhost/health
	http.HandleFunc("/health", s.healthHandler)

	// Runtime memory debugger
	// Read notes/memory_profile.txt
	if dbgMem {
		// curl http://localhost/debug/pprof/heap > heap.0.pprof
		// go tool pprof heap.0.pprof
		// > top10
		http.HandleFunc("/debug/mem/", pprof.Index)
	}
}
