// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/utils/crypto"
	"net/http"
)

// ######################### ACCOUNT API SWITCH #########################

func (s *Service) switchAPIOn(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		DbgPrint("Access Denied! Wrong or missing auth key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	s.state.apiOn = true

	http.StatusText(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! API service switched on! "))
	return
}

func (s *Service) switchAPIOff(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		DbgPrint("Access Denied! Wrong or missing auth key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	s.state.apiOn = false

	http.StatusText(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! API service switched off! "))
	return
}

// ######################### ORDER API SWITCH #########################

func (s *Service) switchOrderOn(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		DbgPrint(accessDeniedError)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	s.state.orderOn = true

	http.StatusText(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! Order service switched on! "))
	return
}

func (s *Service) switchOrderOff(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		DbgPrint(accessDeniedError)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	s.state.orderOn = false

	http.StatusText(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! Order service switched off! "))

	return
}

func (s *Service) resetAllOrderMapHandler(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		DbgPrint(accessDeniedError)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	s.state.orderManager.ResetAllOrderMap()
	http.StatusText(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! All Order Cache Reset! "))
	return
}

func (s *Service) resetLimitOrderMapHandler(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		DbgPrint(accessDeniedError)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	s.state.orderManager.ResetLimitOrderMap()
	http.StatusText(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! Limit Order Cache Reset! "))
	return
}

func (s *Service) resetBookOrderMapHandler(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetOrderAuthKey()) {
		DbgPrint(accessDeniedError)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	s.state.orderManager.ResetBookOrderMap()
	http.StatusText(http.StatusAccepted)
	_, _ = w.Write([]byte("OK! Book Order Cache Reset! "))
	return
}
