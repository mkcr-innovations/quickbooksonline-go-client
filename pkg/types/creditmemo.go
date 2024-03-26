package types

type CreditMemo struct {
	BaseEntity
	Line                    []Line                       `json:"Line"`                              // Required
	CustomerRef             Ref                          `json:"CustomerRef"`                       // Required
	CurrencyRef             *Ref                         `json:"CurrencyRef,omitempty"`             // Conditionally required
	BillEmail               *EmailAddress                `json:"BillEmail,omitempty"`               // Conditionally required
	TxnDate                 string                      `json:"TxnDate,omitempty"`                 // Optional, format "YYYY-MM-DD"
	CustomField             []CustomField                `json:"CustomField,omitempty"`             // Optional
	ClassRef                *Ref                         `json:"ClassRef,omitempty"`                // Optional
	PrintStatus             *PrintStatusEnum             `json:"PrintStatus,omitempty"`             // Optional
	SalesTermRef            *Ref                         `json:"SalesTermRef,omitempty"`            // Optional
	GlobalTaxCalculation    *GlobalTaxCalculationEnum    `json:"GlobalTaxCalculation,omitempty"`    // Optional
	TotalAmt                float64                     `json:"TotalAmt,omitempty"`                // Optional
	InvoiceRef              *Ref                         `json:"InvoiceRef,omitempty"`              // Optional, minorVersion: 37
	TransactionLocationType *TransactionLocationTypeEnum `json:"TransactionLocationType,omitempty"` // Optional, minorVersion: 4
	ApplyTaxAfterDiscount   *bool                        `json:"ApplyTaxAfterDiscount,omitempty"`   // Optional
	DocNumber               *string                      `json:"DocNumber,omitempty"`               // Optional
	PrivateNote             *string                      `json:"PrivateNote,omitempty"`             // Optional
	CustomerMemo            *MemoRef                     `json:"CustomerMemo,omitempty"`            // Optional
	ProjectRef              *Ref                         `json:"ProjectRef,omitempty"`              // Optional, minorVersion: 69
	TxnTaxDetail            *TxnTaxDetail                `json:"TxnTaxDetail,omitempty"`            // Optional
	PaymentMethodRef        *Ref                         `json:"PaymentMethodRef,omitempty"`        // Optional, read-only
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`            // Optional
	ShipAddr                *PhysicalAddress             `json:"ShipAddr,omitempty"`                // Optional
	DepartmentRef           *Ref                         `json:"DepartmentRef,omitempty"`           // Optional
	EmailStatus             *EmailStatusEnum             `json:"EmailStatus,omitempty"`             // Optional
	BillAddr                *PhysicalAddress             `json:"BillAddr,omitempty"`                // Optional
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`                // Optional, read-only
	HomeBalance             *float64                     `json:"HomeBalance,omitempty"`             // Read-only, minorVersion: 3
	RemainingCredit         *float64                     `json:"RemainingCredit,omitempty"`         // Read-only
	RecurDataRef            *Ref                         `json:"RecurDataRef,omitempty"`            // Read-only, minorVersion: 52
	TaxExemptionRef         *Ref                         `json:"TaxExemptionRef,omitempty"`         // Read-only, minorVersion: 21
	Balance                 *float64                     `json:"Balance,omitempty"`                 // Read-only
	HomeTotalAmt            *float64                     `json:"HomeTotalAmt,omitempty"`            // Read-only
}

type CreditMemoPaginatedResponse struct {
	BasePaginatedResponse
	CreditMemo []CreditMemo `json:"CreditMemo"`
}

type CreditMemoResponse struct {
	BaseResponse
	CreditMemo CreditMemo `json:"CreditMemo"`
}
