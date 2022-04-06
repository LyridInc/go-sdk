package model

import (
	"strings"
	"time"
)

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
	ID          string   `json:"id"  binding:"required"`
	AppId       string   `json:"appId"  binding:"required"`
	Name        string   `json:"name"  binding:"required"`
	Language    string   `json:"language" binding:"required"`
	Web         string   `json:"web"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`

	CreatedBy    string    `json:"createdBy" binding:"required"`
	LastActivity time.Time `json:"lastActivity"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

type ModuleRevision struct {
	ID       string `json:"id"`
	ModuleID string `json:"moduleId"`
	CodeUri  string `json:"codeUri"`

	Title		string	   `json:"title"`
	CreatedBy    	string    `json:"createdBy"`
	CreationTime 	time.Time `json:"creationTime"`
	IsActive     	bool      `json:"isActive"`
	IsLastKnown	bool	   `json:"isLastKnown"`

	//Tags []string `json:"tags"`
	SubmitSizeByte int64 `json:"submitSizeByte"`

	Pipeline *StageDefinition `json:"pipeline"`
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

	LastActivity time.Time `json:"lastActivity"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

type FunctionCode struct {
	ID              string    `json:"id"`
	FunctionID      string    `json:"functionId"`
	TargetFramework string    `json:"targetFramework"`
	CreationTime    time.Time `json:"creationTime"`
	CodeUri         string    `json:"codeUri"`
	ImageUri        string    `json:"imageUri"`

	ArtifactSizeByte int64 `json:"artifactSizeByte"`
}

// Definitions (this is used for the user to configure the app/modules/function

type AppDefinition struct {
	Name        string             `yaml:"name"`
	Description string             `yaml:"description"`
	Modules     []ModuleDefinition `yaml:"modules"`
	Volume 		VolumeDefinition 	`yaml:"volume"`
}

type ModuleDefinition struct {
	Name           string               `yaml:"name"`
	Language       string               `yaml:"language"`
	Description    string               `yaml:"description"`
	Web            string               `yaml:"web"`
	ProjectFolder  string               `yaml:"projectFolder"` // currently only used inside dotnet core project only, but technically this works on other languages
	PrebuildScript string               `yaml:"prebuildScript"`
	Functions      []FunctionDefinition `yaml:"functions"`
}

type VolumeDefinition struct {
	Alias 			string				`yaml:"alias"`
	MountPath 		string				`yaml:"mountPath"`
}

type FunctionDefinition struct {
	Name        string `yaml:"name"`
	Entry       string `yaml:"entry"`
	Description string `yaml:"description"`
}

type PublishedApp struct {
	ID          string    `json:"globalId"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Visibility  string    `json:"visibility"`
	SharedWith  []string  `json:"sharedwith"`
	Tier        string    `json:"tier"`
	ImageUrl    string    `json:"imageUrl"`
	SupportUrl  string    `json:"supportUrl"`
	TermUrl     string    `json:"termUrl"`
	TagIds      []string  `json:"tagids"`
	CreatedBy   string    `json:"createdby"`
	CreatedTime time.Time `json:"createTime"`
	LastUpdate  time.Time `json:"lastUpdate"`
}

func (definition *ModuleDefinition) GetFileExtension() string {
	if definition.Language == "go1.x" {
		return "go"
	} else if definition.Language == "python3.7" || definition.Language == "python3.8" || definition.Language == "python3.9" {
		return "py"
	} else if definition.Language == "nodejs12.x" || definition.Language == "nodejs14.x" {
		if strings.Contains(definition.Web, "typescript") {
			return "ts"
		}
		return "js"
	} else if definition.Language == "dotnetcore3.1" || definition.Language == "dotnetcore5.0" {
		return "cs"
	}

	return ""
}
