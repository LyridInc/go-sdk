package model

import "time"

type Tag struct {
	Id          string    `json:"id"`
	AppId       string    `json:"appId"`
	ModuleId    string    `json:"moduleId"`
	AccountId   string    `json:"accountId"`
	RevisionIds []string  `json:"revisionIds"`
	Name        string    `json:"name"`
	CreatedOn   time.Time `json:"createdOn"`
	LastUpdate  time.Time `json:"lastUpdate"`
}
