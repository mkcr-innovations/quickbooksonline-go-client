package types

type JournalCode struct {
	Id          string                `json:"Id"`
	Name        string                `json:"Name"`
	SyncToken   string                `json:"SyncToken"`
	Description *string               `json:"Description,omitempty"`
	CustomField []CustomField         `json:"CustomField,omitempty"`
	Type        *string               `json:"Type,omitempty"`
	MetaData    *ModificationMetaData `json:"MetaData,omitempty"`

	Domain string `json:"domain"`
	Sparse bool   `json:"sparse"`
}

type JournalCodePaginatedResponse struct {
	BasePaginatedResponse
	JournalCode []JournalCode `json:"JournalCode"`
}

type JournalCodeResponse struct {
	BaseResponse
	JournalCode JournalCode `json:"JournalCode"`
}
