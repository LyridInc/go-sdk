package model

import (
	"os"
	"strconv"
)

/*

this is a dynamic query builder that takes in JSON and converts it into BSON query and executes it for a result in the MongoDB.

example query:

{"filters":[{"exp1":"hostname","opt":"$eq","exp2":"vega02.lyrid.io"},{"exp1":"name","opt":"$eq","exp2":"vega01.lyrid.io"}],"Operation":"$or","attribute":{"skip":0,"take":10,"sort":[{"column":"name","direction":"DESC"}]}}

the db and collection for which this query is executed is defined in another place.
*/

type Query struct {
	Filters   []QueryFilter   `json:"filters"`
	Operation string          `json:boolOpt`
	Attribute *QueryAttribute `json:"attribute"`
}

type QueryFilter struct {
	// exp1 is the document name and
	QueryOperand1 string `json:"exp1"`

	// operator can be: >, >=, <, <=, ==, !=
	QueryOperator string `json:"opt"`

	// exp2 is the value usually the value to compare with
	QueryOperand2 interface{} `json:"exp2"`
}

type QueryAttribute struct {
	Skip       int64           `json:"skip"`
	Take       int64           `json:"take"`
	ColumnSort []SortAttribure `json:"sort"`
}

type SortAttribure struct {
	ColumnName string `json:"column"`
	Direction  string `json:"direction"`
}

func NewQueryAttribute(skip int64, take int64, columnSort []SortAttribure) *QueryAttribute {

	max, _ := strconv.Atoi(os.Getenv("MAX_TAKE"))
	if take > int64(max) {
		take = int64(max)
	}

	return &QueryAttribute{Skip: skip, Take: take, ColumnSort: columnSort}
}
