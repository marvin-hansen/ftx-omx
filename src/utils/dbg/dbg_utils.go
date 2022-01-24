// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg

func NilCheck(value interface{}, msg string) {
	if value == nil {
		println(msg)
		panic(value)
	}
}
