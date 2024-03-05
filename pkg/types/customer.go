package types

import (
	"time"
)

// Customer represents a QuickBooks Online customer entity.
type Customer struct {
	Id                      string                `json:"Id"`                                // Read-only, system-defined
	SyncToken               string                `json:"SyncToken"`                         // Read-only, system-defined
	DisplayName             string                `json:"DisplayName"`                       // Conditionally required
	Title                   string                `json:"Title"`                             // Conditionally required, max 16 chars
	GivenName               string                `json:"GivenName"`                         // Conditionally required, max 100 chars
	MiddleName              string                `json:"MiddleName"`                        // Conditionally required, max 100 chars
	Suffix                  string                `json:"Suffix"`                            // Conditionally required, max 16 chars
	FamilyName              string                `json:"FamilyName"`                        // Conditionally required, max 100 chars
	PrimaryEmailAddr        *EmailAddress         `json:"PrimaryEmailAddr,omitempty"`        // Optional
	ResaleNum               *string               `json:"ResaleNum,omitempty"`               // Optional, max 16 chars
	SecondaryTaxIdentifier  *string               `json:"SecondaryTaxIdentifier,omitempty"`  // Optional, minorVersion: 3
	ARAccountRef            *Ref                  `json:"ARAccountRef,omitempty"`            // Optional, minorVersion: 3
	DefaultTaxCodeRef       *Ref                  `json:"DefaultTaxCodeRef,omitempty"`       // Optional
	PreferredDeliveryMethod *string               `json:"PreferredDeliveryMethod,omitempty"` // Optional
	GSTIN                   *string               `json:"GSTIN,omitempty"`                   // Optional, minorVersion: 33
	SalesTermRef            *Ref                  `json:"SalesTermRef,omitempty"`            // Optional
	CustomerTypeRef         *string               `json:"CustomerTypeRef,omitempty"`         // Optional
	Fax                     *TelephoneNumber      `json:"Fax,omitempty"`                     // Optional
	BusinessNumber          *string               `json:"BusinessNumber,omitempty"`          // Optional, minorVersion: 33
	BillWithParent          *bool                 `json:"BillWithParent,omitempty"`          // Optional
	CurrencyRef             *Ref                  `json:"CurrencyRef,omitempty"`             // Optional, read-only
	Mobile                  *TelephoneNumber      `json:"Mobile,omitempty"`                  // Optional
	Job                     *bool                 `json:"Job,omitempty"`                     // Optional
	BalanceWithJobs         *float64              `json:"BalanceWithJobs,omitempty"`         // Optional, sortable
	PrimaryPhone            *TelephoneNumber      `json:"PrimaryPhone,omitempty"`            // Optional
	OpenBalanceDate         *time.Time            `json:"OpenBalanceDate,omitempty"`         // Optional
	Taxable                 *bool                 `json:"Taxable,omitempty"`                 // Optional
	AlternatePhone          *TelephoneNumber      `json:"AlternatePhone,omitempty"`          // Optional
	MetaData                *ModificationMetaData `json:"MetaData,omitempty"`                // Optional, read-only
	ParentRef               *Ref                  `json:"ParentRef,omitempty"`               // Optional
	Notes                   *string               `json:"Notes,omitempty"`                   // Optional, max 2000 chars
	WebAddr                 *WebSiteAddress       `json:"WebAddr,omitempty"`                 // Optional
	Active                  *bool                 `json:"Active,omitempty"`                  // Optional, filterable, sortable
	CompanyName             *string               `json:"CompanyName,omitempty"`             // Optional, max 100 chars
	Balance                 *float64              `json:"Balance,omitempty"`                 // Optional, filterable, sortable
	ShipAddr                *PhysicalAddress      `json:"ShipAddr,omitempty"`                // Optional
	PaymentMethodRef        *Ref                  `json:"PaymentMethodRef,omitempty"`        // Optional
	IsProject               *bool                 `json:"IsProject,omitempty"`               // Optional, read-only, minorVersion: 25
	Source                  *string               `json:"Source,omitempty"`                  // Optional, minorVersion: 59
	PrimaryTaxIdentifier    *string               `json:"PrimaryTaxIdentifier,omitempty"`    // Optional, minorVersion: 4
	GSTRegistrationType     *string               `json:"GSTRegistrationType,omitempty"`     // Optional, minorVersion: 33
	PrintOnCheckName        *string               `json:"PrintOnCheckName,omitempty"`        // Optional, max 110 chars
	BillAddr                *PhysicalAddress      `json:"BillAddr,omitempty"`                // Optional
	FullyQualifiedName      string                `json:"FullyQualifiedName"`                // Read-only, system-defined
	Level                   int                   `json:"Level"`                             // Read-only, system-defined
	TaxExemptionReasonId    *int                  `json:"TaxExemptionReasonId,omitempty"`    // Optional, minorVersion: 10
}

type CustomerPaginatedResponse struct {
	BasePaginatedResponse
	Customer []Customer `json:"Customer"`
}

type CustomerResponse struct {
	BaseResponse
	Customer Customer `json:"Customer"`
}
