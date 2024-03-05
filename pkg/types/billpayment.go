package types

type BillPayment struct {
	Id                      string                 `json:"Id"` // Read only, system defined
	VendorRef               Ref                    `json:"VendorRef"`
	Lines                   []Line                 `json:"Lines"`
	TotalAmt                float64                `json:"TotalAmt"`
	PayType                 string                 `json:"PayType"`
	SyncToken               string                 `json:"SyncToken"`             // Read only, system defined
	CurrencyRef             *Ref                   `json:"CurrencyRef,omitempty"` // Conditionally required
	DocNumber               *string                `json:"DocNumber,omitempty"`
	PrivateNote             *string                `json:"PrivateNote,omitempty"`
	TxnDate                 *string                `json:"TxnDate,omitempty"` // Use string for "YYYY-MM-DD"
	ExchangeRate            *float64               `json:"ExchangeRate,omitempty"`
	APAccountRef            *Ref                   `json:"APAccountRef,omitempty"`
	DepartmentRef           *Ref                   `json:"DepartmentRef,omitempty"`
	TransactionLocationType *string                `json:"TransactionLocationType,omitempty"` // Applies to specific minor version
	ProcessBillPayment      *bool                  `json:"ProcessBillPayment,omitempty"`
	MetaData                *ModificationMetaData  `json:"MetaData,omitempty"`
	CheckPayment            *BillPaymentCheck      `json:"CheckPayment,omitempty"`
	CreditCardPayment       *BillPaymentCreditCard `json:"CreditCardPayment,omitempty"`
}

type BillPaymentCheck struct {
	BankAccountRef Ref     `json:"BankAccountRef"`
	PrintStatus    *string `json:"PrintStatus,omitempty"`
}

type BillPaymentCreditCard struct {
	CCAccountRef Ref `json:"CCAccountRef"`
}

type BillPaymentPaginatedResponse struct {
	BasePaginatedResponse
	BillPayment []BillPayment `json:"BillPayment"`
}

type BillPaymentResponse struct {
	BaseResponse
	BillPayment BillPayment `json:"BillPayment"`
}