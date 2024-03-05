package types

import (
	"time"
)

type Report struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

type ReportHeader struct {
	Customer           string                 `json:"Customer"`
	ReportName         string                 `json:"ReportName"`
	Vendor             string                 `json:"Vendor"`
	Option             []ReportOption         `json:"Option"`
	Item               string                 `json:"Item"`
	Employee           string                 `json:"Employee"`
	ReportBasis        ReportBasisEnum        `json:"ReportBasis"`
	StartPeriod        string                 `json:"StartPeriod"`
	Class              string                 `json:"Class"`
	Currency           string                 `json:"Currency"`
	EndPeriod          string                 `json:"EndPeriod"`
	Time               time.Time              `json:"Time"`
	Department         string                 `json:"Department"`
	SummarizeColumnsBy SummarizeColumnsByEnum `json:"SummarizeColumnsBy"`
}

type ReportOption struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type ReportColumns struct {
	Column []ReportColumn `json:"Column"`
}

type ReportRows struct {
	Row []ReportRow `json:"Row"`
}

type ReportRow struct {
	Type    string                `json:"type"`
	Group   *string               `json:"group,omitempty"`
	ColData []ReportColData       `json:"ColData,omitempty"`
	Summary *ReportColDataWrapper `json:"Summary,omitempty"`
	Rows    *ReportRows           `json:"Rows,omitempty"`
	Header  *ReportColDataWrapper `json:"Header,omitempty"`
}

type ReportColData struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type ReportColDataWrapper struct {
	ColData []ReportColData `json:"ColData"`
}

type ReportColumn struct {
	ColType  ColumnTypeEnum         `json:"ColType"`
	ColTitle *string                `json:"ColTitle,omitempty"`
	MetaData []ReportColumnMetadata `json:"MetaData,omitempty"`
}

type ReportColumnMetadata struct {
	Name  *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// Enum types
type ReportBasisEnum string
type SummarizeColumnsByEnum string
type ColumnTypeEnum string

// Enum values
const (
	ReportBasisCash    ReportBasisEnum = "Cash"
	ReportBasisAccrual ReportBasisEnum = "Accrual"

	SummarizeColumnsByTotal SummarizeColumnsByEnum = "Total"

	ColumnTypeAccount ColumnTypeEnum = "Account"
	ColumnTypeMoney   ColumnTypeEnum = "Money"
)
