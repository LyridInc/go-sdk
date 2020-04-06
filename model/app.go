package model

import "time"

type App struct {
	ID        string `json:"id" binding:"required"`
	AccountId string `json:"accountId" binding:"required"`
	Name      string `json:"name" binding:"required"`
	CreatedBy string `json:"createdBy" binding:"required"`

	Description  string    `json:"description"`
	LastActivity time.Time `json:"lastActivity"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

type Module struct {
	ID        string `json:"id"  binding:"required"`
	AppId     string `json:"appId"  binding:"required"`
	Name      string `json:"name"  binding:"required"`
	Language  string `json:"language" binding:"required"`
	CreatedBy string `json:"createdBy" binding:"required"`

	Description  string    `json:"description"`
	LastActivity time.Time `json:"lastActivity""`
	LastUpdate   time.Time `json:"lastUpdate""`
}

type ModuleRevision struct {
	ID       string `json:"id"`
	ModuleID string `json:"moduleId"`
	CodeUri  string `json:"codeUri"`

	CreatedBy    string    `json:"createdBy"`
	CreationTime time.Time `json:"creationTime"`
	IsActive     bool      `json:"isActive"`
	//Tags []string `json:"tags"`
}

type ModuleBuild struct {
	RevisionID      string    `json:"revisionId"`
	TargetFramework string    `json:"targetFramework"`
	CreationTime    time.Time `json:"creationTime"`
	Uri             string    `json:"uri"`
}

type Function struct {
	ID          string `json:"id"`
	ModuleID    string `json:"moduleId"`
	RevisionID  string `json:"revisionId"`
	Name        string `json:"name"`
	Description string `json:"description"`

	LastActivity time.Time `json:"lastActivity""`
	LastUpdate   time.Time `json:"lastUpdate""`
}

type FunctionCode struct {
	ID              string    `json:"id"`
	FunctionID      string    `json:"functionId"`
	TargetFramework string    `json:"targetFramework"`
	CreationTime    time.Time `json:"creationTime"`
	CodeUri         string    `json:"codeUri"`
	ImageUri        string    `json:"imageUri"`
}
