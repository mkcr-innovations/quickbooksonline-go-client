package types

type PurchaseOrder struct {
	BaseEntity
	APAccountRef            Ref                          `json:"APAccountRef"`
	VendorRef               Ref                          `json:"VendorRef"`
	Line                    []Line                       `json:"Line"`
	CurrencyRef             *Ref                         `json:"CurrencyRef,omitempty"` // Conditionally required, treated as optional.
	TxnDate                 string                      `json:"TxnDate,omitempty"`
	CustomField             []CustomField                `json:"CustomField,omitempty"`
	POEmail                 *EmailAddress                `json:"POEmail,omitempty"` // Applies to minorVersion 17
	ClassRef                *Ref                         `json:"ClassRef,omitempty"`
	SalesTermRef            *Ref                         `json:"SalesTermRef,omitempty"`
	LinkedTxn               []LinkedTxn                  `json:"LinkedTxn,omitempty"`
	GlobalTaxCalculation    *GlobalTaxCalculationEnum    `json:"GlobalTaxCalculation,omitempty"`
	Memo                    *string                      `json:"Memo,omitempty"`
	POStatus                *string                      `json:"POStatus,omitempty"`
	TransactionLocationType *TransactionLocationTypeEnum `json:"TransactionLocationType,omitempty"` // Applies to minorVersion 4
	DueDate                 *string                      `json:"DueDate,omitempty"`
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`
	DocNumber               *string                      `json:"DocNumber,omitempty"`
	PrivateNote             *string                      `json:"PrivateNote,omitempty"`
	ShipMethodRef           *Ref                         `json:"ShipMethodRef,omitempty"`
	TxnTaxDetail            *TxnTaxDetail                `json:"TxnTaxDetail,omitempty"`
	ShipTo                  *Ref                         `json:"ShipTo,omitempty"`
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`
	ShipAddr                *PhysicalAddress             `json:"ShipAddr,omitempty"`
	VendorAddr              *PhysicalAddress             `json:"VendorAddr,omitempty"`
	EmailStatus             *EmailStatusEnum             `json:"EmailStatus,omitempty"`  // Applies to minorVersion 45
	TotalAmt                float64                      `json:"TotalAmt"`               // Read only, system defined
	RecurDataRef            *Ref                         `json:"RecurDataRef,omitempty"` // Applies to minorVersion 52, Read only
}

type PurchaseOrderPaginatedResponse struct {
	BasePaginatedResponse
	PurchaseOrder []PurchaseOrder `json:"PurchaseOrder"`
}

type PurchaseOrderResponse struct {
	BaseResponse
	PurchaseOrder PurchaseOrder `json:"PurchaseOrder"`
}
