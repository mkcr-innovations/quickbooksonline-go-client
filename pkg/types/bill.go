package types

type Bill struct {
	BaseEntity
	VendorRef               Ref                          `json:"VendorRef"`
	Line                    []Line                       `json:"Line"`
	CurrencyRef             Ref                          `json:"CurrencyRef"`
	TxnDate                 string                      `json:"TxnDate,omitempty"`
	APAccountRef            *Ref                         `json:"APAccountRef,omitempty"`
	SalesTermRef            *Ref                         `json:"SalesTermRef,omitempty"`
	LinkedTxn               []LinkedTxn                  `json:"LinkedTxn,omitempty"`
	GlobalTaxCalculation    *GlobalTaxCalculationEnum    `json:"GlobalTaxCalculation,omitempty"`
	TotalAmt                float64                     `json:"TotalAmt,omitempty"`
	TransactionLocationType *TransactionLocationTypeEnum `json:"TransactionLocationType,omitempty"`
	DueDate                 *string                      `json:"DueDate,omitempty"`
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`
	DocNumber               *string                      `json:"DocNumber,omitempty"`
	PrivateNote             *string                      `json:"PrivateNote,omitempty"`
	TxnTaxDetail            *TxnTaxDetail                `json:"TxnTaxDetail,omitempty"`
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Ref                         `json:"DepartmentRef,omitempty"`
	IncludeInAnnualTPAR     *bool                        `json:"IncludeInAnnualTPAR,omitempty"`
	HomeBalance             *float64                     `json:"HomeBalance,omitempty"`
	RecurDataRef            *Ref                         `json:"RecurDataRef,omitempty"`
	Balance                 float64                      `json:"Balance"`
}

type BillPaginatedResponse struct {
	BasePaginatedResponse
	Bill []Bill `json:"Bill"`
}

type BillResponse struct {
	BaseResponse
	Bill Bill `json:"Bill"`
}
