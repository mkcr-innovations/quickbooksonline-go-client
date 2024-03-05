package types

type Vendor struct {
	BaseEntity
	Title                   *string                          `json:"Title,omitempty"`
	GivenName               *string                          `json:"GivenName,omitempty"`
	MiddleName              *string                          `json:"MiddleName,omitempty"`
	Suffix                  *string                          `json:"Suffix,omitempty"`
	FamilyName              *string                          `json:"FamilyName,omitempty"`
	PrimaryEmailAddr        *EmailAddress                    `json:"PrimaryEmailAddr,omitempty"`
	DisplayName             *string                          `json:"DisplayName,omitempty"`
	OtherContactInfo        []ContactInfo                    `json:"OtherContactInfo,omitempty"`
	APAccountRef            *Ref                             `json:"APAccountRef,omitempty"`
	TermRef                 *Ref                             `json:"TermRef,omitempty"`
	Source                  *SourceEnum                      `json:"Source,omitempty"`
	GSTIN                   *string                          `json:"GSTIN,omitempty"`
	T4AEligible             *bool                            `json:"T4AEligible,omitempty"`
	Fax                     *TelephoneNumber                 `json:"Fax,omitempty"`
	BusinessNumber          *string                          `json:"BusinessNumber,omitempty"`
	CurrencyRef             *Ref                             `json:"CurrencyRef,omitempty"`
	HasTPAR                 *bool                            `json:"HasTPAR,omitempty"`
	TaxReportingBasis       *TaxReportingBasisEnum           `json:"TaxReportingBasis,omitempty"`
	Mobile                  *TelephoneNumber                 `json:"Mobile,omitempty"`
	PrimaryPhone            *TelephoneNumber                 `json:"PrimaryPhone,omitempty"`
	Active                  *bool                            `json:"Active,omitempty"`
	AlternatePhone          *TelephoneNumber                 `json:"AlternatePhone,omitempty"`
	MetaData                *ModificationMetaData            `json:"MetaData,omitempty"`
	Vendor1099              *bool                            `json:"Vendor1099,omitempty"`
	CostRate                *float64                         `json:"CostRate,omitempty"`
	BillRate                *float64                         `json:"BillRate,omitempty"`
	WebAddr                 *WebSiteAddress                  `json:"WebAddr,omitempty"`
	T5018Eligible           *bool                            `json:"T5018Eligible,omitempty"`
	CompanyName             *string                          `json:"CompanyName,omitempty"`
	VendorPaymentBankDetail *VendorPaymentBankDetail         `json:"VendorPaymentBankDetail,omitempty"`
	TaxIdentifier           *string                          `json:"TaxIdentifier,omitempty"`
	AcctNum                 *string                          `json:"AcctNum,omitempty"`
	GSTRegistrationType     *CustomerGSTRegistrationTypeEnum `json:"GSTRegistrationType,omitempty"`
	PrintOnCheckName        *string                          `json:"PrintOnCheckName,omitempty"`
	BillAddr                *PhysicalAddress                 `json:"BillAddr,omitempty"`
	Balance                 float64                          `json:"Balance"`
}

type ContactInfo struct {
	Type      *string          `json:"Type,omitempty"`
	Telephone *TelephoneNumber `json:"Telephone,omitempty"`
}

type VendorPaymentBankDetail struct {
	BankAccountName      string  `json:"BankAccountName"`
	BankBranchIdentifier string  `json:"BankBranchIdentifier"`
	BankAccountNumber    string  `json:"BankAccountNumber"`
	StatementText        *string `json:"StatementText,omitempty"`
}

type VendorPaginatedResponse struct {
	BasePaginatedResponse
	Vendor []Vendor `json:"Vendor"`
}

type VendorResponse struct {
	BaseResponse
	Vendor Vendor `json:"Vendor"`
}

type TaxReportingBasisEnum string

const (
	CashTaxReportingBasis    TaxReportingBasisEnum = "Cash"
	AccrualTaxReportingBasis TaxReportingBasisEnum = "Accrual"
)
