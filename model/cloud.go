package model

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
