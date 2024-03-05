package types

type Class struct {
	BaseEntity
	Name               string                `json:"Name"`
	ParentRef          *Ref                  `json:"ParentRef,omitempty"`
	SubClass           *bool                 `json:"SubClass,omitempty"`
	Active             *bool                 `json:"Active,omitempty"`
	MetaData           *ModificationMetaData `json:"MetaData,omitempty"`
	FullyQualifiedName string                `json:"FullyQualifiedName"`
}

type ClassPaginatedResponse struct {
	BasePaginatedResponse
	Class []Class `json:"Class"`
}

type ClassResponse struct {
	BaseResponse
	Class Class `json:"Class"`
}
