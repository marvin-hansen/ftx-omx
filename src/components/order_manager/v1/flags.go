// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"ftx-omx/src/utils/dbg"
)

const (
	debug  = cfg.DbgOrderManager // debug switches non-error handling debugging on & off through the central dbg config file.
	dbgErr = true                // dbgErr switches error handling debugging on & off. True by default to log any error.
	main   = "OrderManager: "    // main component name added in front of each dbg print statement
)

func DbgPrint(msg string) {
	dbg.DbgPrint(debug, main+msg)
}
