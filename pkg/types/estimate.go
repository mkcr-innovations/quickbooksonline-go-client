package types

type Estimate struct {
	BaseEntity
	CustomerRef             Ref                          `json:"CustomerRef"`
	ShipFromAddr            *PhysicalAddress             `json:"ShipFromAddr,omitempty"`            // Conditionally required, minorVersion: 35
	CurrencyRef             Ref                          `json:"CurrencyRef"`                       // Conditionally required
	BillEmail               EmailAddress                 `json:"BillEmail"`                         // Conditionally required
	TxnDate                 string                      `json:"TxnDate,omitempty"`                 // Optional
	ShipDate                *string                      `json:"ShipDate,omitempty"`                // Optional, format: YYYY-MM-DD
	ClassRef                *Ref                         `json:"ClassRef,omitempty"`                // Optional
	PrintStatus             *PrintStatusEnum             `json:"PrintStatus,omitempty"`             // Optional
	CustomField             []CustomField                `json:"CustomField,omitempty"`             // Optional
	SalesTermRef            *Ref                         `json:"SalesTermRef,omitempty"`            // Optional
	TxnStatus               *TxnStatusEnum               `json:"TxnStatus,omitempty"`               // Optional
	LinkedTxn               []LinkedTxn                  `json:"LinkedTxn,omitempty"`               // Optional
	GlobalTaxCalculation    *GlobalTaxCalculationEnum    `json:"GlobalTaxCalculation,omitempty"`    // Optional
	AcceptedDate            *string                      `json:"AcceptedDate,omitempty"`            // Optional, format: YYYY-MM-DD
	ExpirationDate          *string                      `json:"ExpirationDate,omitempty"`          // Optional, format: YYYY-MM-DD
	TransactionLocationType *TransactionLocationTypeEnum `json:"TransactionLocationType,omitempty"` // Optional, minorVersion: 4
	DueDate                 *string                      `json:"DueDate,omitempty"`                 // Optional, format: YYYY-MM-DD
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`                // Optional
	DocNumber               *string                      `json:"DocNumber,omitempty"`               // Optional, max character: 21
	PrivateNote             *string                      `json:"PrivateNote,omitempty"`             // Optional, max character: 4000
	Line                    []Line                       `json:"Line,omitempty"`                    // Optional
	CustomerMemo            *MemoRef                     `json:"CustomerMemo,omitempty"`            // Optional
	EmailStatus             *EmailStatusEnum             `json:"EmailStatus,omitempty"`             // Optional
	ProjectRef              *Ref                         `json:"ProjectRef,omitempty"`              // Optional, minorVersion: 69
	TxnTaxDetail            *TxnTaxDetail                `json:"TxnTaxDetail,omitempty"`            // Optional
	AcceptedBy              *string                      `json:"AcceptedBy,omitempty"`              // Optional
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`            // Optional
	ShipAddr                *PhysicalAddress             `json:"ShipAddr,omitempty"`                // Optional
	DepartmentRef           *Ref                         `json:"DepartmentRef,omitempty"`           // Optional
	ShipMethodRef           *Ref                         `json:"ShipMethodRef,omitempty"`           // Optional
	BillAddr                *PhysicalAddress             `json:"BillAddr,omitempty"`                // Optional
	ApplyTaxAfterDiscount   *bool                        `json:"ApplyTaxAfterDiscount,omitempty"`   // Optional
	TotalAmt                *float64                     `json:"TotalAmt,omitempty"`                // Read only, System defined
	RecurDataRef            *Ref                         `json:"RecurDataRef,omitempty"`            // Read only, minorVersion: 52
	TaxExemptionRef         *Ref                         `json:"TaxExemptionRef,omitempty"`         // Read only, minorVersion: 21
	HomeTotalAmt            *float64                     `json:"HomeTotalAmt,omitempty"`            // Read only, System defined
	FreeFormAddress         *bool                        `json:"FreeFormAddress,omitempty"`         // System defined
}

type EstimatePaginatedResponse struct {
	BasePaginatedResponse
	Estimate []Estimate `json:"Estimate"`
}

type EstimateResponse struct {
	BaseResponse
	Estimate Estimate `json:"Estimate"`
}

type TxnStatusEnum string

const (
	AcceptedTxnStatus  TxnStatusEnum = "Accepted"
	ClosedTxnStatus    TxnStatusEnum = "Closed"
	PendingTxnStatus   TxnStatusEnum = "Pending"
	RejectedTxnStatus  TxnStatusEnum = "Rejected"
	ConvertedTxnStatus TxnStatusEnum = "Converted"
)
