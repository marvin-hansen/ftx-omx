// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"net/http"
)

func (s *Service) rootHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Header
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! 202"))
}

func (s *Service) startHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Header
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Started!"))
}

func (s *Service) statusHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Header
	if s.state.ready {
		_, _ = w.Write([]byte("Online!"))
	} else {
		_, _ = w.Write([]byte("Offline!"))
	}
}

func (s *Service) healthHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Header
	if s.state.ready {
		_, _ = w.Write([]byte("Online"))
	} else { // https://stackoverflow.com/questions/40096750/how-to-set-http-status-code-on-http-responsewriter
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - SERVICE INIT ERROR!"))
	}
}
