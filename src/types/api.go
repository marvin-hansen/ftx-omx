// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import "fmt"

// Api - Our struct for all api's
type Api struct {
	PK          int64  `pg:",pk,unique" json:",omitempty"` // PK for internal DB use
	Id          string `json:"id"`
	AccountName string `json:"accountName"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
}

func (o Api) String() string {
	return fmt.Sprintf("[API]: \n Api ID: %v, \n AccountName: %v, \n API Key: %v, \n API Secret: %v,",
		o.Id,
		o.AccountName,
		o.Key,
		o.Secret,
	)
}
