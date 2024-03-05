package types

type CustomerType struct {
	Id        string                `json:"Id"`
	SyncToken string                `json:"SyncToken"`
	Name      string                `json:"Name"`
	Active    *bool                 `json:"Active,omitempty"`
	MetaData  *ModificationMetaData `json:"MetaData,omitempty"`

	Domain string `json:"domain"`
	Sparse bool   `json:"sparse"`
}

type CustomerTypePaginatedResponse struct {
	BasePaginatedResponse
	CustomerType []CustomerType `json:"CustomerType"`
}

type CustomerTypeResponse struct {
	BaseResponse
	CustomerType CustomerType `json:"CustomerType"`
}
