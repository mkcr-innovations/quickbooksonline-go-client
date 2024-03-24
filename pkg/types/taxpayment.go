package types

type TaxPayment struct {
	BaseEntity
	MetaData          *ModificationMetaData `json:"MetaData,omitempty"`
	Refund            bool                  `json:"Refund,omitempty"`
	TxnDate           string                `json:"TxnDate,omitempty"`
	PaymentAccountRef Ref                   `json:"PaymentAccountRef,omitempty"`
	Description       string                `json:"Description,omitempty"`
	PaymentAmount     float64               `json:"PaymentAmount,omitempty"`
}

type TaxPaymentPaginatedResponse struct {
	BasePaginatedResponse
	TaxPayment []TaxPayment `json:"TaxPayment"`
}

type TaxPaymentResponse struct {
	BaseResponse
	TaxPayment TaxPayment `json:"TaxPayment"`
}
