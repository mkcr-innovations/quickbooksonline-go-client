package types

// Employee represents a QuickBooks Online Employee entity.
type Employee struct {
	BaseEntity
	PrimaryAddr      *PhysicalAddress      `json:"PrimaryAddr,omitempty"` // Conditionally required
	PrimaryEmailAddr *EmailAddress         `json:"PrimaryEmailAddr,omitempty"`
	DisplayName      *string               `json:"DisplayName,omitempty"` // Optional, max 500 chars
	Title            *string               `json:"Title,omitempty"`       // Optional, max 16 chars
	BillableTime     *bool                 `json:"BillableTime,omitempty"`
	GivenName        *string               `json:"GivenName,omitempty"` // Optional, max 100 chars, filterable, sortable
	BirthDate        *string               `json:"BirthDate,omitempty"`
	MiddleName       *string               `json:"MiddleName,omitempty"` // Optional, max 100 chars, filterable, sortable
	SSN              *string               `json:"SSN,omitempty"`        // Optional, max 100 chars
	PrimaryPhone     *TelephoneNumber      `json:"PrimaryPhone,omitempty"`
	Active           *bool                 `json:"Active,omitempty"`
	ReleasedDate     *string               `json:"ReleasedDate,omitempty"`
	MetaData         *ModificationMetaData `json:"MetaData,omitempty"`
	CostRate         *float64              `json:"CostRate,omitempty"` // BigDecimal
	Mobile           *TelephoneNumber      `json:"Mobile,omitempty"`
	Gender           *GenderEnum           `json:"Gender,omitempty"`
	HiredDate        *string               `json:"HiredDate,omitempty"`
	BillRate         *float64              `json:"BillRate,omitempty"` // BigDecimal, optional
	Organization     *bool                 `json:"Organization,omitempty"`
	Suffix           *string               `json:"Suffix,omitempty"`           // Optional, max 16 chars
	FamilyName       *string               `json:"FamilyName,omitempty"`       // Optional, max 100 chars, filterable, sortable
	PrintOnCheckName *string               `json:"PrintOnCheckName,omitempty"` // Optional, max 100 chars
	EmployeeNumber   *string               `json:"EmployeeNumber,omitempty"`   // Optional, max 100 chars
	V4IDPseudonym    *string               `json:"V4IDPseudonym,omitempty"`    // minorVersion: 26, read only
}

type EmployeePaginatedResponse struct {
	BasePaginatedResponse
	Employee []Employee `json:"Employee"`
}

type EmployeeResponse struct {
	BaseResponse
	Employee Employee `json:"Employee"`
}

type GenderEnum string

const (
	MaleGender   GenderEnum = "Male"
	FemaleGender GenderEnum = "Female"
)
