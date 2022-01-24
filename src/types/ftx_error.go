// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import "time"

type Error struct {
	PK     int64     `pg:",pk,unique" json:",omitempty"` // PK for internal DB use
	Time   time.Time `json:"time"`
	Market string    `json:"market"`
	Error  string    `json:"error"`
}

func NewError(market string, error string) *Error {
	return &Error{
		Time:   time.Now(),
		Market: market,
		Error:  error,
	}
}
