package model

import (
	"strings"
	"time"
)

type App struct {
	ID          string `json:"id" binding:"required"`
	AccountId   string `json:"accountId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	RelatedVega string `json:"relatedVega"`
	Namespace   string `json:"namespace"`
	Alias       string `json:"alias"`
	CreatedBy   string `json:"createdBy" binding:"required"`

	Description  string    `json:"description"`
	LastActivity time.Time `json:"lastActivity"`
	LastUpdate   time.Time `json:"lastUpdate"`

	GitURL string `json:"gitUrl"`

	DistributedRegion bool `json:"distributedRegion"`
	IsDeleted         bool `json:"isDeleted"`
	UseOperator       bool `json:"useOperator"`
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

	Title        string    `json:"title"`
	CreatedBy    string    `json:"createdBy"`
	CreationTime time.Time `json:"creationTime"`
	IsActive     bool      `json:"isActive"`
	IsLastKnown  bool      `json:"isLastKnown"`

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
	Name          string                  `yaml:"name"`
	Description   string                  `yaml:"description"`
	IgnoreFiles   string                  `yaml:"ignoreFiles"`
	Modules       []ModuleDefinition      `yaml:"modules"`
	Database      DatabaseDefinition      `yaml:"database"`
	ObjectStorage ObjectStorageDefinition `yaml:"objectStorage"`
	Resources     []ResourcesDefinition   `yaml:"resources"`
	Spec          []SpecDefinition        `yaml:"spec"`
}

type ResourcesDefinition struct {
	Cpu      string       `yaml:"cpu"`
	Memory   string       `yaml:"memory"`
	Requests ResourceList `yaml:"requests"`
	Limits   ResourceList `yaml:"limits"`
}

type SpecDefinition struct {
	Replica string `yaml:"replica"`
}

type ResourceList struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type ModuleDefinition struct {
	ID             string               `json:"id"`
	Name           string               `yaml:"name"`
	Language       string               `yaml:"language"`
	Description    string               `yaml:"description"`
	Web            string               `yaml:"web"`
	HidePublicURL  bool                 `yaml:"hidePublicURL"`
	ProjectFolder  string               `yaml:"projectFolder"`
	PrebuildScript string               `yaml:"prebuildScript"`
	Config         ConfigDefinition     `yaml:"config"`
	Functions      []FunctionDefinition `yaml:"functions"`

	Volumes      []VolumeDefinition  `yaml:"volumes"`
	Ports        []PortDefinition    `yaml:"ports"`
	Resources    ResourcesDefinition `json:"resources"`
	CustomLabels []KVPairStandard    `yaml:"customLabels"`

	LastActivity time.Time `json:"lastActivity"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

type KVPairStandard struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type VolumeDefinition struct {
	Alias     string `yaml:"alias"`
	MountPath string `yaml:"mountPath"`
}

type PortDefinition struct {
	Alias    string `yaml:"alias"`
	Port     int64  `yaml:"port"`
	UseProbe bool   `yaml:"useProbe"`
}

type ConfigDefinition struct {
	Instance    string           `yaml:"instance"`
	RegionId    string           `yaml:"regionId"`
	Distributed bool             `yaml:"distributed"`
	Scale       *AutoScaleConfig `yaml:"scale"`
	PublicURL   *bool            `yaml:"publicURL"`
}

type AutoScaleConfig struct {
	Min int64 `yaml:"min"`
	Max int64 `yaml:"max"`
}

type FunctionDefinition struct {
	Name        string `yaml:"name"`
	Entry       string `yaml:"entry"`
	Description string `yaml:"description"`
}

type DatabaseDefinition struct {
	Alias string `yaml:"alias"`
	Type  string `yaml:"type"`
}

type ObjectStorageDefinition struct {
	Alias string `yaml:"alias"`
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
	// languages := map[string][]string{
	// 	"go": []string{"go1.x"},
	// 	"py": []string{"python3.7", "python3.8", "python3.9", "python3.10", "python3.11"},
	// 	"ts": [],
	// }

	if definition.Language == "go1.x" {
		return "go"
	} else if definition.Language == "python3.7" || definition.Language == "python3.8" || definition.Language == "python3.9" || definition.Language == "python3.10" || definition.Language == "python3.11" {
		return "py"
	} else if definition.Language == "nodejs12.x" || definition.Language == "nodejs14.x" || definition.Language == "nodejs16.x" || definition.Language == "nodejs18.x" || definition.Language == "nodejs20.x" {
		if strings.Contains(definition.Web, "typescript") {
			return "ts"
		}
		return "js"
	} else if definition.Language == "dotnetcore3.1" || definition.Language == "dotnetcore5.0" || definition.Language == "dotnetcore7.0" {
		return "cs"
	}

	return ""
}
