package model

import (
	"strings"
	"time"
)

type LyridUser struct {
	Id              string    `json:"id" binding:"required"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerified   bool      `json:"emailVerified"`
	Roles           []string  `json:"roles"`
	RelatedRoles    []string  `json:"relatedRoles`
	RelatedAccounts []string  `json:"relatedAccounts`
	BetaAccess      bool      `json:"betaAccess"`
	CurrentAccount  string    `json:"currentAccount"`
	DefaultAccount  string    `json:"defaultAccount"`
	LastUpdate      time.Time `json:"lastUpdate"`
}

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Tier      int       `json:"tier" binding:"required"`
	CreatedOn time.Time `json:"createdOn"`
	CreatedBy string    `json:"createdBy"`
}

func (account *Account) GetBucketName() string {
	return strings.ReplaceAll(account.Id, "-", "")
}

func (account *Account) GetS3BucketName(region string) string {
	return "lyrid-lambda-" + account.GetBucketName() + "-" + region
}
