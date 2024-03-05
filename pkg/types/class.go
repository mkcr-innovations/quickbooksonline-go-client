package types

type Class struct {
	Id                 string                `json:"Id"`
	Name               string                `json:"Name"`
	SyncToken          string                `json:"SyncToken"`
	ParentRef          *Ref                  `json:"ParentRef,omitempty"`
	SubClass           *bool                 `json:"SubClass,omitempty"`
	Active             *bool                 `json:"Active,omitempty"`
	MetaData           *ModificationMetaData `json:"MetaData,omitempty"`
	FullyQualifiedName string                `json:"FullyQualifiedName"`

	Domain string `json:"domain"`
	Sparse bool   `json:"sparse"`
}

type ClassPaginatedResponse struct {
	BasePaginatedResponse
	Class []Class `json:"Class"`
}

type ClassResponse struct {
	BaseResponse
	Class Class `json:"Class"`
}
