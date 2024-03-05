package types

type Deposit struct {
	Id                      string                `json:"Id"` // Read Only, System Defined
	DepositToAccountRef     Ref                   `json:"DepositToAccountRef"`
	Lines                   []Line                `json:"Line"`                              // Required
	SyncToken               *string               `json:"SyncToken,omitempty"`               // Read Only, System Defined
	CurrencyRef             *Ref                  `json:"CurrencyRef,omitempty"`             // Conditionally Required
	PrivateNote             *string               `json:"PrivateNote,omitempty"`             // Optional, Max character: max of 4000 chars
	GlobalTaxCalculation    *string               `json:"GlobalTaxCalculation,omitempty"`    // Optional, minorVersion: 3
	ExchangeRate            *float64              `json:"ExchangeRate,omitempty"`            // Optional
	DepartmentRef           *Ref                  `json:"DepartmentRef,omitempty"`           // Optional
	TxnSource               *string               `json:"TxnSource,omitempty"`               // Optional
	TxnDate                 *string               `json:"TxnDate,omitempty"`                 // Optional, Date
	CashBack                *CashBackInfo         `json:"CashBack,omitempty"`                // Optional
	TransactionLocationType *string               `json:"TransactionLocationType,omitempty"` // Optional, minorVersion: 4
	TxnTaxDetail            *TxnTaxDetail         `json:"TxnTaxDetail,omitempty"`            // Optional, minorVersion: 4
	MetaData                *ModificationMetaData `json:"MetaData,omitempty"`                // Optional
	RecurDataRef            *Ref                  `json:"RecurDataRef,omitempty"`            // Read Only, minorVersion: 52
	TotalAmt                *float64              `json:"TotalAmt,omitempty"`                // Read Only, System Defined
	HomeTotalAmt            *float64              `json:"HomeTotalAmt,omitempty"`            // Read Only, System Defined
}

type CashBackInfo struct {
	AccountRef Ref     `json:"AccountRef"`     // Required
	Amount     float64 `json:"Amount"`         // Required
	Memo       *string `json:"Memo,omitempty"` // Optional
}

type DepositPaginatedResponse struct {
	BasePaginatedResponse
	Deposit []Deposit `json:"Deposit"`
}

type DepositResponse struct {
	BaseResponse
	Deposit Deposit `json:"Deposit"`
}
