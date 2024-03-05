package types

type Transfer struct {
	Id                      string                `json:"Id"`                                // Required for update, read only, system defined
	ToAccountRef            Ref                   `json:"ToAccountRef"`                      // Required
	Amount                  float64               `json:"Amount"`                            // Required
	FromAccountRef          Ref                   `json:"FromAccountRef"`                    // Required
	SyncToken               string                `json:"SyncToken"`                         // Required for update, read only, system defined
	PrivateNote             *string               `json:"PrivateNote,omitempty"`             // Optional, Max of 4000 chars
	TxnDate                 *string               `json:"TxnDate,omitempty"`                 // Optional, Date
	TransactionLocationType *string               `json:"TransactionLocationType,omitempty"` // Optional, minorVersion: 4
	MetaData                *ModificationMetaData `json:"MetaData,omitempty"`                // Optional
	RecurDataRef            *Ref                  `json:"RecurDataRef,omitempty"`            // Read only, minorVersion: 52
}

type TransferPaginatedResponse struct {
	BasePaginatedResponse
	Transfer []Transfer `json:"Transfer"`
}

type TransferResponse struct {
	BaseResponse
	Transfer Transfer `json:"Transfer"`
}
