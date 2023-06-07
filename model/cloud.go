package model

import (
	"time"

	"github.com/alecthomas/units"
)

type CloudVendor int

const (
	AWS CloudVendor = 1 + iota
	GCP
	LYR
)

var vendors = [...]string{
	"AWS",
	"GCP",
	"LYR",
}

func (c CloudVendor) String() string { return vendors[c-1] }

type CloudVendorDefinition struct {
	VendorID      CloudVendor `json:"vendorId" binding:"required"`
	Name          string      `json:"name" binding:"required"`
	ShortName     string      `json:"shortName" binding:"required"`
	ParentCompany string      `json:"parentCompany" binding:"required"`
	ImageUrl      string      `json:"imageUrl" binding:"required"`
}

type GlobalRegionDefinition struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`

	Description string `json:"description"`

	Center *LngLatGeo `json:"center"`
	Radius int        `json:"radius"`

	RegionIDs []string `json:"regionIds"`
}

type CloudCredential struct {
	Id             string      `json:"id" binding:"required"`
	KeyId          string      `json:"keyId" binding:"required"`
	VendorID       CloudVendor `json:"vendorId" binding:"required"`
	RelatedAccount string      `json:"relatedAccount" binding:"required"`
	CreationTime   time.Time   `json:"creationTime" binding:"required"`
	UseCount       int         `json:"useCount" binding:"required"`
	Credential     []byte      `json:"credential" binding:"required"`
}

type DeploymentEndpoint struct {
	ID string `json:"id" binding:"required"`

	CodeID  string   `json:"codeId" binding:"required"`
	CodeIDs []string `json:"codeIds"`
	Name    string   `json:"name" binding:"required"`

	Type     string      `json:"type" binding:"required"`
	VendorID CloudVendor `json:"vendorId" binding:"required"`
	RegionID string      `json:"regionId" binding:"required"`

	Endpoint string        `json:"endpoint" binding:"required"`
	Memory   string        `json:"memory" binding:"required"`
	Timeout  time.Duration `json:"duration" binding:"required"`

	Metadata  string `json:"metadata"`
	Namespace string `json:"namespace"`

	RelatedVega string `json:"relatedvega"`
}

type FrameworkDefinition struct {
	Name     string `json:"name" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`
}

type StorageDefinition struct {
	Name     string `json:"name" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`

	Type     string `json:"type" binding:"required"`
	Endpoint string `json:"endpoint" binding:"required"`
}

type CloudRegionDefinition struct {
	ID string `json:"id" binding:"required"`

	VendorID CloudVendor `json:"vendorId" binding:"required"`

	RegionID string `json:"regionId" binding:"required"`
	Type     string `json:"type" binding:"required"`

	Location *LngLatGeo `json:"location" binding:"required"`

	Framework FrameworkDefinition `json:"framework" binding:"required"`
	Storage   StorageDefinition   `json:"storage" binding:"required"`
	//
}

type LngLatGeo struct {
	Longitute float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`

	Address *LocationDescription `json:"address"`
}

type LocationDescription struct {
	City          AddressName `json:"city"`
	Country       AddressName `json:"country"`
	StateProvince AddressName `json:"stateProvince"`
}

type AddressName struct {
	LongName  string `json:"longName"`
	ShortName string `json:"shortName"`
}

func (deployment *DeploymentEndpoint) GetInitialConsumptionUnit() int64 {
	b, _ := units.ParseBase2Bytes(deployment.Memory + "B")
	return int64(b) / int64(128*units.MiB)
}
