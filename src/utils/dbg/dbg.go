// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg

import (
	"log"
)

func DbgCheck(dbg bool) {
	if dbg {
		log.Println("Debug mode on!")
	}
}

func DbgPrint(dbg bool, msg string) {
	if dbg {
		log.Println(msg)
	}
}
