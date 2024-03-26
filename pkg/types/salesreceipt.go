package types

type SalesReceipt struct {
	BaseEntity
	Line                    []Line                       `json:"Line"`                              // Required
	CustomerRef             Ref                          `json:"CustomerRef"`                       // Required
	ShipFromAddr            *PhysicalAddress             `json:"ShipFromAddr,omitempty"`            // Conditionally required, minorVersion: 35
	CurrencyRef             *Ref                         `json:"CurrencyRef,omitempty"`             // Conditionally required
	BillEmail               *EmailAddress                `json:"BillEmail,omitempty"`               // Conditionally required
	TxnDate                 string                      `json:"TxnDate,omitempty"`                 // Optional
	CustomField             []CustomField                `json:"CustomField,omitempty"`             // Optional
	ShipDate                *string                      `json:"ShipDate,omitempty"`                // Optional
	TrackingNum             *string                      `json:"TrackingNum,omitempty"`             // Optional
	ClassRef                *Ref                         `json:"ClassRef,omitempty"`                // Optional
	PrintStatus             *PrintStatusEnum             `json:"PrintStatus,omitempty"`             // Optional
	PaymentRefNum           *string                      `json:"PaymentRefNum,omitempty"`           // Optional, Max 21 characters
	TxnSource               *string                      `json:"TxnSource,omitempty"`               // Optional
	LinkedTxn               []LinkedTxn                  `json:"LinkedTxn,omitempty"`               // Optional
	GlobalTaxCalculation    *GlobalTaxCalculationEnum    `json:"GlobalTaxCalculation,omitempty"`    // Optional
	TransactionLocationType *TransactionLocationTypeEnum `json:"TransactionLocationType,omitempty"` // Optional, minorVersion: 4
	ApplyTaxAfterDiscount   *bool                        `json:"ApplyTaxAfterDiscount,omitempty"`   // Optional
	DocNumber               *string                      `json:"DocNumber,omitempty"`               // Optional, Max 21 characters
	PrivateNote             *string                      `json:"PrivateNote,omitempty"`             // Optional, Max 4000 characters
	DepositToAccountRef     *Ref                         `json:"DepositToAccountRef,omitempty"`     // Optional
	CustomerMemo            *MemoRef                     `json:"CustomerMemo,omitempty"`            // Optional
	EmailStatus             *EmailStatusEnum             `json:"EmailStatus,omitempty"`             // Optional
	ProjectRef              *Ref                         `json:"ProjectRef,omitempty"`              // Optional, minorVersion: 69
	CreditCardPayment       *CreditCardPaymentForPayment `json:"CreditCardPayment,omitempty"`       // Optional
	TxnTaxDetail            *TxnTaxDetail                `json:"TxnTaxDetail,omitempty"`            // Optional
	PaymentMethodRef        *Ref                         `json:"PaymentMethodRef,omitempty"`        // Optional
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`            // Optional
	ShipAddr                *PhysicalAddress             `json:"ShipAddr,omitempty"`                // Optional
	DepartmentRef           *Ref                         `json:"DepartmentRef,omitempty"`           // Optional
	ShipMethodRef           *Ref                         `json:"ShipMethodRef,omitempty"`           // Optional
	BillAddr                *PhysicalAddress             `json:"BillAddr,omitempty"`                // Optional
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`                // Optional
	HomeBalance             *float64                     `json:"HomeBalance,omitempty"`             // ReadOnly, minorVersion: 3
	DeliveryInfo            DeliveryInfo                 `json:"DeliveryInfo"`                      // ReadOnly
	RecurDataRef            *Ref                         `json:"RecurDataRef,omitempty"`            // ReadOnly, minorVersion: 52
	TotalAmt                float64                      `json:"TotalAmt"`                          // ReadOnly
	Balance                 float64                      `json:"Balance"`                           // ReadOnly
	HomeTotalAmt            float64                      `json:"HomeTotalAmt"`                      // ReadOnly
	FreeFormAddress         bool                         `json:"FreeFormAddress"`                   // System defined
}

type SalesReceiptPaginatedResponse struct {
	BasePaginatedResponse
	SalesReceipt []SalesReceipt `json:"SalesReceipt"`
}

type SalesReceiptResponse struct {
	BaseResponse
	SalesReceipt SalesReceipt `json:"SalesReceipt"`
}
