package types

type Department struct {
	Id                 string                `json:"Id"`
	Name               string                `json:"Name"`
	SyncToken          string                `json:"SyncToken"`
	ParentRef          *Ref                  `json:"ParentRef,omitempty"`
	Active             *bool                 `json:"Active,omitempty"`
	MetaData           *ModificationMetaData `json:"MetaData,omitempty"`
	FullyQualifiedName string                `json:"FullyQualifiedName"`
	SubDepartment      bool                  `json:"SubDepartment"`

	Domain string `json:"domain"`
	Sparse bool   `json:"sparse"`
}

type DepartmentPaginatedResponse struct {
	BasePaginatedResponse
	Department []Department `json:"Department"`
}

type DepartmentResponse struct {
	BaseResponse
	Department Department `json:"Department"`
}
