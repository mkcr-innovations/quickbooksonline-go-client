package types

import (
	"time"
)

type Budget struct {
	BaseEntity
	EndDate         time.Time             `json:"EndDate"`
	StartDate       time.Time             `json:"StartDate"`
	BudgetEntryType *BudgetEntryTypeEnum  `json:"BudgetEntryType,omitempty"` // Optional
	Name            *string               `json:"Name,omitempty"`            // Optional
	BudgetDetail    []BudgetDetail        `json:"BudgetDetail,omitempty"`    // Optional
	BudgetType      *BudgetTypeEnum       `json:"BudgetType,omitempty"`      // Optional
	Active          *bool                 `json:"Active,omitempty"`          // Optional
	MetaData        *ModificationMetaData `json:"MetaData,omitempty"`        // Optional
}

type BudgetDetail struct {
	ClassRef      *Ref       `json:"ClassRef,omitempty"`      // Optional
	DepartmentRef *Ref       `json:"DepartmentRef,omitempty"` // Optional
	Amount        *float64   `json:"Amount,omitempty"`        // Optional
	BudgetDate    *time.Time `json:"BudgetDate,omitempty"`    // Optional
	AccountRef    *Ref       `json:"AccountRef,omitempty"`    // Optional
	CustomerRef   *Ref       `json:"CustomerRef,omitempty"`   // Optional
}

type BudgetPaginatedResponse struct {
	BasePaginatedResponse
	Budget []Budget `json:"Budget"`
}

type BudgetResponse struct {
	BaseResponse
	Budget Budget `json:"Budget"`
}

type BudgetEntryTypeEnum string

const (
	MonthlyBudgetEntryType   BudgetEntryTypeEnum = "Monthly"
	QuarterlyBudgetEntryType BudgetEntryTypeEnum = "Quarterly"
	AnnuallyBudgetEntryType  BudgetEntryTypeEnum = "Annually"
)

type BudgetTypeEnum string

const ProfitAndLossBudgetType BudgetTypeEnum = "ProfitAndLoss"
