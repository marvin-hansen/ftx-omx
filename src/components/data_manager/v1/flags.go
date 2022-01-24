// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/utils/dbg"
)

const (
	debug = cfg.DbgDataManager
	main  = "[DataManger] "
)

func DbgPrint(mtd, msg string) {
	dbg.DbgPrint(debug, main+mtd+msg)
}
