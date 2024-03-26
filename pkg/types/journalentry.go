package types

type JournalEntry struct {
	BaseEntity
	Line                    []Line                `json:"Line"`                              // Required
	CurrencyRef             *Ref                  `json:"CurrencyRef,omitempty"`             // Conditionally required
	DocNumber               *string               `json:"DocNumber,omitempty"`               // Optional
	PrivateNote             *string               `json:"PrivateNote,omitempty"`             // Optional
	TxnDate                 string               `json:"TxnDate,omitempty"`                 // Optional, use string for "YYYY-MM-DD"
	ExchangeRate            *float64              `json:"ExchangeRate,omitempty"`            // Optional
	TaxRateRef              *Ref                  `json:"TaxRateRef,omitempty"`              // Optional, minorVersion: 49
	TransactionLocationType *TransactionLocationTypeEnum               `json:"TransactionLocationType,omitempty"` // Optional, minorVersion: 4
	TxnTaxDetail            *TxnTaxDetail         `json:"TxnTaxDetail,omitempty"`            // Optional
	GlobalTaxCalculation    *GlobalTaxCalculationEnum               `json:"GlobalTaxCalculation,omitempty"`    // Optional, minorVersion: 53
	Adjustment              *bool                 `json:"Adjustment,omitempty"`              // Optional
	MetaData                *ModificationMetaData `json:"MetaData,omitempty"`                // Optional
	RecurDataRef            *Ref                  `json:"RecurDataRef,omitempty"`            // Read only, minorVersion: 52
	TotalAmt                float64               `json:"TotalAmt"`                          // Read only, system defined, always 0
	HomeTotalAmt            float64               `json:"HomeTotalAmt"`                      // Read only, system defined, always 0
}

type JournalEntryPaginatedResponse struct {
	BasePaginatedResponse
	JournalEntry []JournalEntry `json:"JournalEntry"`
}

type JournalEntryResponse struct {
	BaseResponse
	JournalEntry JournalEntry `json:"JournalEntry"`
}
