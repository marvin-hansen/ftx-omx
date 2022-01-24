// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package account_manager

import (
	"github.com/go-numb/go-ftx/rest"
	v1 "web_socket/src/components/account_manager/v1"
	"web_socket/src/components/api_manager"
)

type AccountManager interface {
	// GetAccount returns the account for the client & ticker.
	// Ensure to defer client close to prevent memory leak i.e.
	// account, clErr := accountManager.GetAccount(clientID, ticker)
	// defer account.HTTPC.CloseIdleConnections()
	// if clErr != nil { ... }
	GetAccount(clientID, ticker string) (c *rest.Client, err error)
	SetLeverage(clientID, ticker string, leverage int) (ok bool, msg string)
	ResetLeverage(clientID, ticker string) (ok bool, msg string)
}

func NewAccountManager(apiManager api_manager.ApiManager) AccountManager {
	return v1.NewAccountManager(apiManager)

}
