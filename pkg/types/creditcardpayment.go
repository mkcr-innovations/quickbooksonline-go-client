package types

type CreditCardPayment struct {
	BaseEntity
	CreditCardAccountRef Ref                   `json:"CreditCardAccountRef"`  // Required
	Amount               float64               `json:"Amount"`                // Required
	BankAccountRef       Ref                   `json:"BankAccountRef"`        // Required
	PrivateNote          *string               `json:"PrivateNote,omitempty"` // Optional, Max of 4000 chars
	VendorRef            *Ref                  `json:"VendorRef,omitempty"`   // Optional, minorVersion: 54
	TxnDate              string               `json:"TxnDate,omitempty"`     // Optional, Date
	Memo                 *string               `json:"Memo,omitempty"`        // Optional, Max of 4000 chars, minorVersion: 54
	PrintStatus          *PrintStatusEnum      `json:"PrintStatus,omitempty"` // Optional, minorVersion: 54
	CheckNum             *string               `json:"CheckNum,omitempty"`    // Optional, minorVersion: 54
	MetaData             *ModificationMetaData `json:"MetaData,omitempty"`    // Optional
}

type CreditCardPaymentPaginatedResponse struct {
	BasePaginatedResponse
	CreditCardPayment []CreditCardPayment `json:"CreditCardPaymentTxn"`
}

type CreditCardPaymentResponse struct {
	BaseResponse
	CreditCardPayment CreditCardPayment `json:"CreditCardPaymentTxn"`
}
