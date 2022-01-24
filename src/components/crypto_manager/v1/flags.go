// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/utils/dbg"
)

const (
	debug = cfg.DbgCryptoManager
	main  = "CryptoManager: "
)

func DbgPrint(msg string) {
	dbg.DbgPrint(debug, main+msg)
}
