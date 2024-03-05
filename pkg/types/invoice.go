package types

import (
	"time"
)

type Invoice struct {
	BaseEntity
	Line                         []Line                `json:"Line"`
	CustomerRef                  Ref                   `json:"CustomerRef"`
	ShipFromAddr                 *PhysicalAddress      `json:"ShipFromAddr,omitempty"` // Conditionally required, minorVersion: 35
	CurrencyRef                  Ref                   `json:"CurrencyRef"`
	DocNumber                    string                `json:"DocNumber"`
	BillEmail                    EmailAddress          `json:"BillEmail"`
	TxnDate                      *string               `json:"TxnDate,omitempty"`  // Format: YYYY-MM-DD
	ShipDate                     *string               `json:"ShipDate,omitempty"` // Format: YYYY-MM-DD
	TrackingNum                  *string               `json:"TrackingNum,omitempty"`
	ClassRef                     *Ref                  `json:"ClassRef,omitempty"`
	PrintStatus                  *string               `json:"PrintStatus,omitempty"`
	SalesTermRef                 *Ref                  `json:"SalesTermRef,omitempty"`
	TxnSource                    *string               `json:"TxnSource,omitempty"`
	LinkedTxn                    []LinkedTxn           `json:"LinkedTxn,omitempty"`
	DepositToAccountRef          *Ref                  `json:"DepositToAccountRef,omitempty"`
	GlobalTaxCalculation         *string               `json:"GlobalTaxCalculation,omitempty"`
	AllowOnlineACHPayment        *bool                 `json:"AllowOnlineACHPayment,omitempty"`
	TransactionLocationType      *string               `json:"TransactionLocationType,omitempty"` // minorVersion: 4
	DueDate                      *string               `json:"DueDate,omitempty"`                 // Format: YYYY-MM-DD
	MetaData                     *ModificationMetaData `json:"MetaData,omitempty"`
	PrivateNote                  *string               `json:"PrivateNote,omitempty"` // max character: 4000
	BillEmailCc                  *EmailAddress         `json:"BillEmailCc,omitempty"` // minorVersion: 8
	CustomerMemo                 *MemoRef              `json:"CustomerMemo,omitempty"`
	EmailStatus                  *string               `json:"EmailStatus,omitempty"`
	ProjectRef                   *Ref                  `json:"ProjectRef,omitempty"` // minorVersion: 69
	ExchangeRate                 *float64              `json:"ExchangeRate,omitempty"`
	Deposit                      *float64              `json:"Deposit,omitempty"`
	TxnTaxDetail                 *TxnTaxDetail         `json:"TxnTaxDetail,omitempty"`
	AllowOnlineCreditCardPayment *bool                 `json:"AllowOnlineCreditCardPayment,omitempty"`
	CustomField                  []CustomField         `json:"CustomField,omitempty"`
	ShipAddr                     *PhysicalAddress      `json:"ShipAddr,omitempty"`
	DepartmentRef                *Ref                  `json:"DepartmentRef,omitempty"`
	BillEmailBcc                 *EmailAddress         `json:"BillEmailBcc,omitempty"` // minorVersion: 8
	ShipMethodRef                *Ref                  `json:"ShipMethodRef,omitempty"`
	BillAddr                     *PhysicalAddress      `json:"BillAddr,omitempty"`
	ApplyTaxAfterDiscount        *bool                 `json:"ApplyTaxAfterDiscount,omitempty"`
	HomeBalance                  *float64              `json:"HomeBalance,omitempty"`     // Read only, minorVersion: 3
	DeliveryInfo                 *DeliveryInfo         `json:"DeliveryInfo,omitempty"`    // Read only
	TotalAmt                     *float64              `json:"TotalAmt,omitempty"`        // Read only, System defined
	InvoiceLink                  *string               `json:"InvoiceLink,omitempty"`     // Read only, minorVersion: 36, System defined
	RecurDataRef                 *Ref                  `json:"RecurDataRef,omitempty"`    // Read only, minorVersion: 52
	TaxExemptionRef              *Ref                  `json:"TaxExemptionRef,omitempty"` // Read only, minorVersion: 21, System defined
	Balance                      float64               `json:"Balance"`                   // Read only
	HomeTotalAmt                 float64               `json:"HomeTotalAmt"`              // Read only
	FreeFormAddress              bool                  `json:"FreeFormAddress"`           // System defined
	// Deprecated fields
	AllowOnlinePayment *bool `json:"AllowOnlinePayment,omitempty"` // Deprecated
	AllowIPNPayment    *bool `json:"AllowIPNPayment,omitempty"`    // Deprecated
}

type DeliveryInfo struct {
	DeliveryType string    `json:"DeliveryType"` // Read only
	DeliveryTime time.Time `json:"DeliveryTime"` // Read only
}

type InvoicePaginatedResponse struct {
	BasePaginatedResponse
	Invoice []Invoice `json:"Invoice"`
}

type InvoiceResponse struct {
	BaseResponse
	Invoice Invoice `json:"Invoice"`
}
