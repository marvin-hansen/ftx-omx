// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package types

type Env int

const (
	UnknownEnv Env = iota
	Dev
	Test
	Prod
)
