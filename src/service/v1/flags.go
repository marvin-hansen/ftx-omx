// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/utils/dbg"
)

const (
	main    = "Service/main: "
	dbgMain = cfg.DbgMain
	dbgMem  = cfg.DbgMemory
)

func DbgPrint(msg string) {
	dbg.DbgPrint(dbgMain, main+msg)
}
