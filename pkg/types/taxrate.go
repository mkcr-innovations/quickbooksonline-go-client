package types

import (
	"time"
)

type TaxRate struct {
	BaseEntity
	RateValue        *float64              `json:"RateValue,omitempty"`        // Optional, ReadOnly
	Name             *string               `json:"Name,omitempty"`             // Optional, ReadOnly, MaxChars: 100
	AgencyRef        *Ref                  `json:"AgencyRef,omitempty"`        // Optional, ReadOnly
	SpecialTaxType   *string               `json:"SpecialTaxType,omitempty"`   // Optional, ReadOnly
	EffectiveTaxRate []EffectiveTaxRate    `json:"EffectiveTaxRate,omitempty"` // Optional, ReadOnly
	DisplayType      *string               `json:"DisplayType,omitempty"`      // Optional, ReadOnly
	TaxReturnLineRef *Ref                  `json:"TaxReturnLineRef,omitempty"` // Optional, ReadOnly
	Active           *bool                 `json:"Active,omitempty"`           // Optional, ReadOnly
	MetaData         *ModificationMetaData `json:"MetaData,omitempty"`         // Optional
	OriginalTaxRate  *string               `json:"OriginalTaxRate,omitempty"`  // Optional, ReadOnly, Applies to minorVersion: 62
	Description      *string               `json:"Description,omitempty"`      // Optional, ReadOnly, MaxChars: 100
}

type EffectiveTaxRate struct {
	RateValue     float64   `json:"RateValue"`     // Decimal
	EndDate       time.Time `json:"EndDate"`       // String, Format: YYYY-MM-DDTHH:MM:SS+/-HH:MM
	EffectiveDate time.Time `json:"EffectiveDate"` // String, Format: YYYY-MM-DDTHH:MM:SS+/-HH:MM
}

type TaxRatePaginatedResponse struct {
	BasePaginatedResponse
	TaxRate []TaxRate `json:"TaxRate"`
}

type TaxRateResponse struct {
	BaseResponse
	TaxRate TaxRate `json:"TaxRate"`
}
