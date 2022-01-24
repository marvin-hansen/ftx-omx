// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/cfg"
	"web_socket/src/utils/dbg"
)

const (
	debug = cfg.DbgClientManager
	main  = "AccountManager: "
)

func DbgPrint(msg string) {
	dbg.DbgPrint(debug, main+msg)
}
