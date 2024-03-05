package types

type CompanyInfo struct {
	Id                        string                `json:"Id"`
	SyncToken                 string                `json:"SyncToken"`
	CompanyName               string                `json:"CompanyName"`
	CompanyAddr               Address               `json:"CompanyAddr"`
	LegalAddr                 *Address              `json:"LegalAddr,omitempty"`
	SupportedLanguages        *string               `json:"SupportedLanguages,omitempty"`
	Country                   *string               `json:"Country,omitempty"`
	Email                     *Email                `json:"Email,omitempty"`
	WebAddr                   *WebAddr              `json:"WebAddr,omitempty"`
	NameValue                 []NameValue           `json:"NameValue,omitempty"`
	FiscalYearStartMonth      *string               `json:"FiscalYearStartMonth,omitempty"`
	CustomerCommunicationAddr *Address              `json:"CustomerCommunicationAddr,omitempty"`
	PrimaryPhone              *PrimaryPhone         `json:"PrimaryPhone,omitempty"`
	LegalName                 *string               `json:"LegalName,omitempty"`
	MetaData                  *ModificationMetaData `json:"MetaData,omitempty"`
	CompanyStartDate          string                `json:"CompanyStartDate"`

	Domain string `json:"domain"`
	Sparse bool   `json:"sparse"`
}

type Address struct {
	Id                     string  `json:"Id"`
	PostalCode             *string `json:"PostalCode,omitempty"`
	City                   *string `json:"City,omitempty"`
	Country                *string `json:"Country,omitempty"`
	Line5                  *string `json:"Line5,omitempty"`
	Line4                  *string `json:"Line4,omitempty"`
	Line3                  *string `json:"Line3,omitempty"`
	Line2                  *string `json:"Line2,omitempty"`
	Line1                  *string `json:"Line1,omitempty"`
	Lat                    *string `json:"Lat,omitempty"`
	Long                   *string `json:"Long,omitempty"`
	CountrySubDivisionCode *string `json:"CountrySubDivisionCode,omitempty"`
}

type WebAddr struct {
	URI *string `json:"URI,omitempty"`
}

type PrimaryPhone struct {
	FreeFormNumber *string `json:"FreeFormNumber,omitempty"`
}

type Email struct {
	Address *string `json:"Address,omitempty"`
}

type NameValue struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type CompanyInfoPaginatedResponse struct {
	BasePaginatedResponse
	CompanyInfo []CompanyInfo `json:"CompanyInfo"`
}

type CompanyInfoResponse struct {
	BaseResponse
	CompanyInfo CompanyInfo `json:"CompanyInfo"`
}
