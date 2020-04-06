package model

import "time"

type UserAccessToken struct {
	Selected     string
	Key          string    `json:"key" binding:"required"`
	Secret       string    `json:"secret" binding:"required"`
	Token        string    `json:"token" binding:"required"`
	IsValid      bool      `json:"isValid" binding:"required"`
	CreationTime time.Time `json:"creationTime"`
	LastActivity time.Time `json:"lastActivity"`
}
