package types

type VendorCredit struct {
	BaseEntity
	VendorRef               Ref                          `json:"VendorRef"`                         // Required
	Lines                   []Line                       `json:"Lines"`                             // Required
	CurrencyRef             *Ref                         `json:"CurrencyRef,omitempty"`             // Conditionally required, thus pointer and omitempty
	DocNumber               *string                      `json:"DocNumber,omitempty"`               // Optional
	PrivateNote             *string                      `json:"PrivateNote,omitempty"`             // Optional
	LinkedTxn               []LinkedTxn                  `json:"LinkedTxn,omitempty"`               // Optional, applies to specific minorversion
	GlobalTaxCalculation    *GlobalTaxCalculationEnum    `json:"GlobalTaxCalculation,omitempty"`    // Optional
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`            // Optional
	APAccountRef            *Ref                         `json:"APAccountRef,omitempty"`            // Optional
	DepartmentRef           *Ref                         `json:"DepartmentRef,omitempty"`           // Optional
	TxnDate                 string                      `json:"TxnDate,omitempty"`                 // Optional, using string for date "YYYY-MM-DD"
	IncludeInAnnualTPAR     *bool                        `json:"IncludeInAnnualTPAR,omitempty"`     // Optional, applies to specific minorversion
	TransactionLocationType *TransactionLocationTypeEnum `json:"TransactionLocationType,omitempty"` // Optional, applies to specific minorversion
	Balance                 *float64                     `json:"Balance,omitempty"`                 // Optional, read-only, applies to specific minorversion
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`                // Optional
	RecurDataRef            *Ref                         `json:"RecurDataRef,omitempty"`            // Optional, read-only, applies to specific minorversion
	TotalAmt                float64                      `json:"TotalAmt"`                          // Optional, read-only, system defined
}

type VendorCreditPaginatedResponse struct {
	BasePaginatedResponse
	VendorCredit []VendorCredit `json:"VendorCredit"`
}

type VendorCreditResponse struct {
	BaseResponse
	VendorCredit VendorCredit `json:"VendorCredit"`
}
