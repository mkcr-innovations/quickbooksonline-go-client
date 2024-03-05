package types

type Department struct {
	BaseEntity
	Name               string                `json:"Name"`
	ParentRef          *Ref                  `json:"ParentRef,omitempty"`
	Active             *bool                 `json:"Active,omitempty"`
	MetaData           *ModificationMetaData `json:"MetaData,omitempty"`
	FullyQualifiedName string                `json:"FullyQualifiedName"`
	SubDepartment      bool                  `json:"SubDepartment"`
}

type DepartmentPaginatedResponse struct {
	BasePaginatedResponse
	Department []Department `json:"Department"`
}

type DepartmentResponse struct {
	BaseResponse
	Department Department `json:"Department"`
}
