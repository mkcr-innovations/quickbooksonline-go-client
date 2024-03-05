package types

type Purchase struct {
	BaseEntity
	Line                    []Line                `json:"Line"`                              // Required
	PaymentType             string                `json:"PaymentType"`                       // Required
	AccountRef              Ref                   `json:"AccountRef"`                        // Required
	CurrencyRef             *Ref                  `json:"CurrencyRef,omitempty"`             // Conditionally required
	TxnDate                 *string               `json:"TxnDate,omitempty"`                 // Optional, format YYYY-MM-DD
	PrintStatus             *string               `json:"PrintStatus,omitempty"`             // Optional
	RemitToAddr             *PhysicalAddress      `json:"RemitToAddr,omitempty"`             // Optional, read only
	TxnSource               *string               `json:"TxnSource,omitempty"`               // Optional
	LinkedTxn               []LinkedTxn           `json:"LinkedTxn,omitempty"`               // Optional, minorVersion: 55
	GlobalTaxCalculation    *string               `json:"GlobalTaxCalculation,omitempty"`    // Optional
	TransactionLocationType *string               `json:"TransactionLocationType,omitempty"` // Optional, minorVersion: 4
	MetaData                *ModificationMetaData `json:"MetaData,omitempty"`                // Optional
	DocNumber               *string               `json:"DocNumber,omitempty"`               // Optional, max 21 chars
	PrivateNote             *string               `json:"PrivateNote,omitempty"`             // Optional, max 4000 chars
	Credit                  *bool                 `json:"Credit,omitempty"`                  // Optional
	TxnTaxDetail            *TxnTaxDetail         `json:"TxnTaxDetail,omitempty"`            // Optional
	PaymentMethodRef        *Ref                  `json:"PaymentMethodRef,omitempty"`        // Optional
	ExchangeRate            *float64              `json:"ExchangeRate,omitempty"`            // Optional
	DepartmentRef           *Ref                  `json:"DepartmentRef,omitempty"`           // Optional
	EntityRef               *Ref                  `json:"EntityRef,omitempty"`               // Optional
	IncludeInAnnualTPAR     *bool                 `json:"IncludeInAnnualTPAR,omitempty"`     // Optional, minorVersion: 40
	TotalAmt                *float64              `json:"TotalAmt,omitempty"`                // Read only, system defined
	RecurDataRef            *Ref                  `json:"RecurDataRef,omitempty"`            // Read only, minorVersion: 52
}

type PurchasePaginatedResponse struct {
	BasePaginatedResponse
	Purchase []Purchase `json:"Purchase"`
}

type PurchaseResponse struct {
	BaseResponse
	Purchase Purchase `json:"Purchase"`
}
