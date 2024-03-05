package types

type ReimburseCharge struct {
	BaseEntity
	Lines           []Line                `json:"Line"`                   // Required
	Amount          float64               `json:"Amount"`                 // Required
	CustomerRef     Ref                   `json:"CustomerRef"`            // Required
	CurrencyRef     *Ref                  `json:"CurrencyRef,omitempty"`  // Conditionally Required
	PrivateNote     *string               `json:"PrivateNote,omitempty"`  // Optional
	LinkedTxn       []LinkedTxn           `json:"LinkedTxn,omitempty"`    // Optional
	ExchangeRate    *float64              `json:"ExchangeRate,omitempty"` // Optional
	MetaData        *ModificationMetaData `json:"MetaData,omitempty"`     // Optional
	HasBeenInvoiced bool                  `json:"HasBeenInvoiced"`
	HomeTotalAmt    *float64              `json:"HomeTotalAmt,omitempty"` // Optional, Read Only, System Defined
}

type ReimburseChargePaginatedResponse struct {
	BasePaginatedResponse
	ReimburseCharge []ReimburseCharge `json:"ReimburseCharge"`
}

type ReimburseChargeResponse struct {
	BaseResponse
	ReimburseCharge ReimburseCharge `json:"ReimburseCharge"`
}
