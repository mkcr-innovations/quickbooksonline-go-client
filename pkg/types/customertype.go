package types

type CustomerType struct {
	BaseEntity
	Name     string                `json:"Name"`
	Active   *bool                 `json:"Active,omitempty"`
	MetaData *ModificationMetaData `json:"MetaData,omitempty"`
}

type CustomerTypePaginatedResponse struct {
	BasePaginatedResponse
	CustomerType []CustomerType `json:"CustomerType"`
}

type CustomerTypeResponse struct {
	BaseResponse
	CustomerType CustomerType `json:"CustomerType"`
}
