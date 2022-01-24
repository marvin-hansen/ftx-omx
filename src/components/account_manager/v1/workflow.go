// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"errors"
	"ftx-omx/src/utils/dbg"
	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/account"
)

func (c AccountManager) GetAccount(clientID, ticker string) (client *rest.Client, err error) {
	c.state.RLock()
	defer c.state.RUnlock()

	client, exists := c.state.clientMap.Get(clientID + ticker)
	if exists {
		DbgPrint("Return account client from map")
		return client, nil

	} else {
		DbgPrint("Create new account client!")
		api, apiErr := c.dependencies.apiManager.GetApi(clientID)
		if apiErr != nil {
			return nil, errors.New("No API found for " + clientID)
		}

		client = rest.New(
			auth.New(
				api.Key,
				api.Secret,
				auth.SubAccount{
					UUID:     1,
					Nickname: api.AccountName,
				},
			))

		DbgPrint("Switch to sub-account")
		client.Auth.UseSubAccountID(1) // or 2... this number is key in map[int]SubAccount
	}

	DbgPrint("Add account client to map")
	c.state.clientMap.Set(clientID+ticker, client)
	return client, nil
}

func (c AccountManager) SetLeverage(clientID, ticker string, leverage int) (ok bool, msg string) {
	if leverage > 1 {
		client, clErr := c.GetAccount(clientID, ticker)
		defer client.HTTPC.CloseIdleConnections()
		if clErr != nil {
			msg = "No API available for: " + clientID
			return false, msg
		}

		_, err := client.Leverage(&account.RequestForLeverage{
			Leverage: leverage,
		})
		if err != nil {
			dbg.LogError(err)
			msg = "FTX server error setting leverage"
			return false, msg
		} else {
			msg = "ok"
			return true, msg
		}
	} else {
		msg = "leverage below one. Abort"
		return false, msg
	}
}

func (c AccountManager) ResetLeverage(clientID, ticker string) (ok bool, msg string) {
	client, clErr := c.GetAccount(clientID, ticker)
	defer client.HTTPC.CloseIdleConnections()
	if clErr != nil {
		msg = "No API available for: " + clientID
		return false, msg
	}

	_, err := client.Leverage(&account.RequestForLeverage{
		Leverage: 1,
	})
	if err != nil {
		dbg.LogError(err)
		msg = "FTX server error resetting leverage"
		return false, msg

	} else {
		msg = "ok"
		return true, msg
	}
}
