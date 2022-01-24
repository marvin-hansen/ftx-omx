// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/types"
	"ftx-omx/src/utils/crypto"
	dbg2 "ftx-omx/src/utils/dbg"
	"log"
	"net/http"
	"strconv"
)

func (s *Service) createApiHandler(w http.ResponseWriter, r *http.Request) {

	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetApiAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint("check if API  service is switched on!")
	if s.state.apiOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(apiServiceUnavailableError))
		return
	}

	DbgPrint("decode JSON request")
	var api types.Api
	err := decodeJSONBody(w, r, &api)
	if err != nil {
		dbg2.LogError(err)
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	DbgPrint("generate new api ID")
	id := generateID(12) // 12 char long ID
	api.Id = id

	DbgPrint(" store api with ID as key: " + api.Id)
	msg = " store api with ID as key: " + api.Id
	DbgPrint(msg)
	s.state.apiManager.AddApi(api)

	DbgPrint(" Monitor api with ID as key: " + api.Id)
	go s.state.apiManager.StartMonitorApi(api)

	msg = "OK! New API ID: " + api.Id + " "
	DbgPrint(msg)
	_, _ = w.Write([]byte(msg))
}

func (s *Service) setApiLeverageHandler(w http.ResponseWriter, r *http.Request) {
	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetApiAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint("check if API  service is switched on!")
	if s.state.apiOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(apiServiceUnavailableError))
		return
	}

	DbgPrint("decode JSON request")
	var lev types.LeverageRequest
	err := decodeJSONBody(w, r, &lev)
	if err != nil {
		dbg2.LogError(err)
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	key := lev.ApiID
	sym := lev.Market
	leverage := lev.Leverage

	if sym == "" {
		msg = "Error: No symbol given!"
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
		return
	}

	if key == "" {
		msg = "Error: No key given!"
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
		return
	}

	exists := s.state.apiManager.CheckApiExists(key)
	if !exists {
		msg = "Warning! API does not exists: " + key
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
		return
	}

	if leverage == 0 {
		msg = "Error: No leverage given! Only use values > 1"
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
		return
	}

	if leverage < 0 {
		msg = "Error: leverage negative! Only use values > 1"
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
		return
	}

	dbg2.DbgLeverageReq(dbgMain, sym, key, leverage)
	ok, errMsg := s.state.clientManager.SetLeverage(key, sym, leverage)
	if ok {
		msg = "OK! Leverage set for API: " + key
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
	} else {
		msg = "Server Error: Failure to set leverage!"
		DbgPrint(msg)
		DbgPrint(errMsg)
		_, _ = w.Write([]byte(msg))
	}
}

func (s *Service) resetApiLeverageHandler(w http.ResponseWriter, r *http.Request) {
	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetApiAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	DbgPrint("check if API  service is switched on!")
	if s.state.apiOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(apiServiceUnavailableError))
		return
	}

	DbgPrint("decode JSON request")
	var lev types.LeverageRequest
	err := decodeJSONBody(w, r, &lev)
	if err != nil {
		dbg2.LogError(err)
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	key := lev.ApiID
	sym := lev.Market

	if sym == "" {
		msg = "Error! No symbol given!"
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
		return
	}

	if key == "" {
		msg = "Error! No key given!"
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
		return
	}

	exists := s.state.apiManager.CheckApiExists(key)
	if !exists {
		msg = "OK! Leverage has been reset! "
		DbgPrint(msg)
		_, _ = w.Write([]byte("Warning! API does not exists: " + key))
		return
	}

	dbg2.DbgLeverageReq(dbgMain, sym, key, 1)
	ok, errMsg := s.state.clientManager.ResetLeverage(key, sym)
	if ok {
		msg = "OK! Leverage has been reset! "
		DbgPrint(msg)
		_, _ = w.Write([]byte("OK! Leverage has been reset! "))
	} else {
		msg = "Server Error: Failure to reset leverage!"
		DbgPrint(msg)
		DbgPrint(errMsg)
		_, _ = w.Write([]byte(msg))
	}
}

func (s *Service) deleteApiHandler(w http.ResponseWriter, r *http.Request) {
	msg := ""
	authKey := r.Header.Get("auth")
	if authKey != crypto.DecodeKey(cfg.GetApiAuthKey()) {
		msg = "API Access denied!"
		DbgPrint(msg + "Incorrect API key!")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		_, _ = w.Write([]byte(msg))
		return
	}

	DbgPrint("check if API  service is switched on!")
	if s.state.apiOn == false {
		msg = "API Disabled"
		DbgPrint(msg)
		http.StatusText(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(msg))
		return
	}

	key := r.FormValue("id")
	if key != "" {
		DbgPrint("extracted key: " + key)
		ok := s.state.apiManager.CheckApiExists(key)
		DbgPrint("Key exists: " + strconv.FormatBool(ok))

		if ok {
			s.state.apiManager.DeleteApi(key)
			msg = "OK! API Deleted: " + key
			DbgPrint(msg)

			DbgPrint(" Stop API Monitor: " + key)
			go s.state.apiManager.StopMonitorApi(key)

			_, _ = w.Write([]byte(msg))
		} else {
			msg = "Warning! API does not exists: " + key
			DbgPrint(msg)
			_, _ = w.Write([]byte(msg))
		}
	} else {
		msg = "Error! No key given!"
		DbgPrint(msg)
		_, _ = w.Write([]byte(msg))
	}
}
