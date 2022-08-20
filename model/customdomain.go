package model

import "time"

type LBTypeEnum uint

const (
	Failover LBTypeEnum = iota
	RoundRobin
	Weighted
)

type CustomDomainStatusEnum uint

const (
	Added CustomDomainStatusEnum = iota
	Valid
)

type CustomDomain struct {
	Id         string             `json:"id" binding:"required"`
	Name       string             `json:"name" binding:"required"`
	AccountId  string             `json:"accountId" binding:"required"`
	Subdomains []SubDomainMapping `json:"subdomains" binding:"required"`
	Redirect   bool               `json:"redirect" binding:"required"`

	LBType LBTypeEnum `json:"lbType" binding:"required"`

	Status CustomDomainStatusEnum `json:"status" binding:"required"`

	Certificates []Certificate `json:"certificates" binding:"required"`
}

type Certificate struct {
	FileName  string    `json:"fileName"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type SubDomainMapping struct {
	Name       string `json:"name"`
	PrefixPath string `json:"prefixPath"`
}
