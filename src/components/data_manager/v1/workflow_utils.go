// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"errors"
	"web_socket/src/types"
)

func (c *DataManager) encryptApi(api types.Api) (encryptedApi types.Api, err error) {
	mtd := "[encryptApi]: "

	DbgPrint(mtd, "Encrypt secret key")
	encryptedKey, msg, ok := c.dependencies.cryptoManager.Encrypt(api.Key)
	if !ok {
		return types.Api{}, errors.New(msg)
	}

	encryptedSecret, msg, ok := c.dependencies.cryptoManager.Encrypt(api.Secret)
	if !ok {
		return types.Api{}, errors.New(msg)
	}

	encryptedAccountName, msg, ok := c.dependencies.cryptoManager.Encrypt(api.AccountName)
	if !ok {
		return types.Api{}, errors.New(msg)
	}

	DbgPrint(mtd, "Constructed encrypted API object ")

	encryptedApi = types.Api{
		Id:          api.Id,
		AccountName: encryptedAccountName,
		Key:         encryptedKey,
		Secret:      encryptedSecret,
	}

	return encryptedApi, nil
}

func (c *DataManager) decryptApi(encryptedApi types.Api) (decryptedApi types.Api, err error) {

	mtd := "decryptApi"

	DbgPrint(mtd, "Decrypt secret key")
	decryptedKey, msg, ok := c.dependencies.cryptoManager.Decrypt(encryptedApi.Key)
	if !ok {
		return types.Api{}, errors.New(msg)
	}

	decryptedSecret, msg, ok := c.dependencies.cryptoManager.Decrypt(encryptedApi.Secret)
	if !ok {
		return types.Api{}, errors.New(msg)
	}

	decryptedAccountName, msg, ok := c.dependencies.cryptoManager.Decrypt(encryptedApi.AccountName)
	if !ok {
		return types.Api{}, errors.New(msg)
	}

	DbgPrint(mtd, "Constructed plain text API object ")
	decryptedApi = types.Api{
		Id:          encryptedApi.Id,
		AccountName: decryptedAccountName,
		Key:         decryptedKey,
		Secret:      decryptedSecret,
	}

	return decryptedApi, err
}
