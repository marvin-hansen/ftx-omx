// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

type InitFunc func(msg string)

func (s *Service) runInit(msg string, initFn InitFunc) {
	initFn(msg)
}
