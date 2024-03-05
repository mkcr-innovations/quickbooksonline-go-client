package types

type RefundReceipt struct {
	BaseEntity
	DepositToAccountRef     Ref                          `json:"DepositToAccountRef"`
	Lines                   []Line                       `json:"Lines"`
	CurrencyRef             *Ref                         `json:"CurrencyRef,omitempty"`   // Conditionally required
	PaymentRefNum           *string                      `json:"PaymentRefNum,omitempty"` // Conditionally required
	BillEmail               *EmailAddress                `json:"BillEmail,omitempty"`     // Conditionally required
	TxnDate                 *string                      `json:"TxnDate,omitempty"`
	CustomField             []CustomField                `json:"CustomField,omitempty"`
	ClassRef                *Ref                         `json:"ClassRef,omitempty"`
	PrintStatus             *PrintStatusEnum             `json:"PrintStatus,omitempty"`
	CheckPayment            *CheckPayment                `json:"CheckPayment,omitempty"`
	TxnSource               *string                      `json:"TxnSource,omitempty"`
	GlobalTaxCalculation    *GlobalTaxCalculationEnum    `json:"GlobalTaxCalculation,omitempty"`
	TransactionLocationType *TransactionLocationTypeEnum `json:"TransactionLocationType,omitempty"` // Applies to minor version 4
	MetaData                *ModificationMetaData        `json:"MetaData,omitempty"`
	DocNumber               *string                      `json:"DocNumber,omitempty"`
	PrivateNote             *string                      `json:"PrivateNote,omitempty"`
	CustomerMemo            *MemoRef                     `json:"CustomerMemo,omitempty"`
	ProjectRef              *Ref                         `json:"ProjectRef,omitempty"` // Applies to minor version 69
	CreditCardPayment       *CreditCardPaymentForPayment `json:"CreditCardPayment,omitempty"`
	CustomerRef             *Ref                         `json:"CustomerRef,omitempty"`
	TxnTaxDetail            *TxnTaxDetail                `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef        *Ref                         `json:"PaymentMethodRef,omitempty"`
	ExchangeRate            *float64                     `json:"ExchangeRate,omitempty"`
	ShipAddr                *PhysicalAddress             `json:"ShipAddr,omitempty"`
	DepartmentRef           *Ref                         `json:"DepartmentRef,omitempty"`
	PaymentType             *PaymentTypeEnum             `json:"PaymentType,omitempty"`
	BillAddr                *PhysicalAddress             `json:"BillAddr,omitempty"`
	ApplyTaxAfterDiscount   *bool                        `json:"ApplyTaxAfterDiscount,omitempty"`
	HomeBalance             *float64                     `json:"HomeBalance,omitempty"`     // Read only, System defined, Applies to minor version 3
	RecurDataRef            *Ref                         `json:"RecurDataRef,omitempty"`    // Read only, System defined, Applies to minor version 52
	TotalAmt                *float64                     `json:"TotalAmt,omitempty"`        // Read only, System defined
	TaxExemptionRef         *Ref                         `json:"TaxExemptionRef,omitempty"` // Read only, System defined, Applies to minor version 21
	Balance                 *float64                     `json:"Balance,omitempty"`         // Read only, System defined
	HomeTotalAmt            *float64                     `json:"HomeTotalAmt,omitempty"`    // Read only, System defined
}

// Define CheckPayment struct as per its child attributes
type CheckPayment struct {
	CheckNum   *string `json:"CheckNum,omitempty"`
	Status     *string `json:"Status,omitempty"`
	NameOnAcct *string `json:"NameOnAcct,omitempty"`
	AcctNum    *string `json:"AcctNum,omitempty"`
	BankName   *string `json:"BankName,omitempty"`
}

type RefundReceiptPaginatedResponse struct {
	BasePaginatedResponse
	RefundReceipt []RefundReceipt `json:"RefundReceipt"`
}

type RefundReceiptResponse struct {
	BaseResponse
	RefundReceipt RefundReceipt `json:"RefundReceipt"`
}
