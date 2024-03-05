package types

type Payment struct {
	BaseEntity
	TotalAmt                float64                      `json:"TotalAmt"`
	CustomerRef             Ref                          `json:"CustomerRef"`
	CurrencyRef             *Ref                         `json:"CurrencyRef,omitempty"` // Conditionally Required
	PrivateNote             *string                      `json:"PrivateNote,omitempty"` // Max of 4000 chars
	PaymentMethodRef        *Ref                         `json:"PaymentMethodRef,omitempty"`
	UnappliedAmt            *float64                     `json:"UnappliedAmt,omitempty"` // ReadOnly
	DepositToAccountRef     *Ref                         `json:"DepositToAccountRef,omitempty"`
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`
	Line                    []Line                       `json:"Line,omitempty"`
	ProjectRef              *Ref                         `json:"ProjectRef,omitempty"` // minorVersion: 69
	TxnSource               *string                      `json:"TxnSource,omitempty"`
	ARAccountRef            *Ref                         `json:"ARAccountRef,omitempty"`
	TxnDate                 *string                      `json:"TxnDate,omitempty"` // Format: YYYY-MM-DD
	CreditCardPayment       *CreditCardPaymentForPayment `json:"CreditCardPayment,omitempty"`
	TransactionLocationType *TransactionLocationTypeEnum                      `json:"TransactionLocationType,omitempty"` // minorVersion: 4
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`
	PaymentRefNum           *string                      `json:"PaymentRefNum,omitempty"`   // Conditionally Required
	TaxExemptionRef         *Ref                         `json:"TaxExemptionRef,omitempty"` // ReadOnly, minorVersion: 21
}

type PaymentPaginatedResponse struct {
	BasePaginatedResponse
	Payment []Payment `json:"Payment"`
}

type PaymentResponse struct {
	BaseResponse
	Payment Payment `json:"Payment"`
}
