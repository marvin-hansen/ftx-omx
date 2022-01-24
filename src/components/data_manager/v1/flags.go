// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"web_socket/src/cfg"
	"web_socket/src/utils/dbg"
)

const (
	debug = cfg.DbgDataManager
	main  = "[DataManger] "
)

func DbgPrint(mtd, msg string) {
	dbg.DbgPrint(debug, main+mtd+msg)
}
