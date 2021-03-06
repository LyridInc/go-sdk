package model

import (
	"strconv"
	"time"
)

type PolicyDefinition struct {
	Id        string `json:"id"`
	AccountID string `json:"accountId"`
	ModuleID  string `json:"moduleId"`

	Policies map[string]string `json:"policies"`

	CreatedOn  time.Time `json:"createdOn"`
	LastUpdate time.Time `json:"lastUpdate"`
}

type Policy struct {
	Id          string `json:"id"`
	LabelName   string `json:"labelName"`
	Description string `json:"description"`

	PolicyType string `json:"policyType"`
	ValueType  string `json:"valueType"`

	DefaultValue interface{} `json:"defaultValue"`
	CurrentValue interface{} `json:"currentValue"`

	PossibleValues []interface{} `json:"possibleValues"`
	Visibility     string        `json:"visibility"`
}

func (policy *Policy) GetValue() interface{} {
	if policy.ValueType == "int" {
		value, _ := strconv.Atoi(policy.CurrentValue.(string))
		return value
	} else if policy.ValueType == "bool" {
		value, _ := strconv.ParseBool(policy.CurrentValue.(string))
		return value
	} else if policy.ValueType == "json" {
		return policy.CurrentValue
	}

	return nil
}

func (policy *Policy) ValidateValue(value string) bool {
	// tbd
	return true
}

func (policy *Policy) SetValue(input string) {
	if policy.ValidateValue(input) {
		policy.CurrentValue = input
	}
}
