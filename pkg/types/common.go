package types

import "time"

// Response types
type BaseResponse struct {
	Time time.Time `json:"time"`
}

type QueryResponse[T any] struct {
	BaseResponse
	QueryResponse T `json:"QueryResponse"`
}

type TotalCountResponse struct {
	TotalCount int `json:"totalCount"`
}

type QueryCountResponse struct {
	QueryResponse[TotalCountResponse]
}

type BasePaginatedResponse struct {
	StartPosition int `json:"startPosition"`
	MaxResults    int `json:"maxResults"`
}

// Base Entity
type BaseEntity struct {
	Id        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
	Domain    string `json:"domain"`
	Sparse    bool   `json:"sparse"`
}

// Common types
type ModificationMetaData struct {
	CreateTime      time.Time `json:"CreateTime"`
	LastUpdatedTime time.Time `json:"LastUpdatedTime"`
}

type Ref struct {
	Name  *string `json:"name,omitempty"`
	Value string  `json:"value"`
}

type MemoRef struct {
	Value string `json:"value"`
}

type PhysicalAddress struct {
	Id                     string  `json:"Id"`
	PostalCode             *string `json:"PostalCode,omitempty"`
	City                   *string `json:"City,omitempty"`
	Country                *string `json:"Country,omitempty"`
	Line1                  *string `json:"Line1,omitempty"`
	Line2                  *string `json:"Line2,omitempty"`
	Line3                  *string `json:"Line3,omitempty"`
	Line4                  *string `json:"Line4,omitempty"`
	Line5                  *string `json:"Line5,omitempty"`
	Lat                    *string `json:"Lat,omitempty"`
	Long                   *string `json:"Long,omitempty"`
	CountrySubDivisionCode *string `json:"CountrySubDivisionCode,omitempty"`
}

type TelephoneNumber struct {
	FreeFormNumber *string `json:"FreeFormNumber,omitempty"`
}

type EmailAddress struct {
	Address *string `json:"Address,omitempty"`
}

type LinkedTxn struct {
	TxnId     string `json:"TxnId"`
	TxnType   string `json:"TxnType"`
	TxnLineId string `json:"TxnLineId"`
}

type TxnTaxDetail struct {
	TxnTaxCodeRef *Ref      `json:"TxnTaxCodeRef,omitempty"`
	TotalTax      *float64  `json:"TotalTax,omitempty"`
	TaxLine       []TaxLine `json:"TaxLine,omitempty"`
}

type TaxLine struct {
	DetailType    string        `json:"DetailType"`
	TaxLineDetail TaxLineDetail `json:"TaxLineDetail"`
	Amount        *float64      `json:"Amount,omitempty"`
}

type TaxLineDetail struct {
	TaxRateRef          Ref      `json:"TaxRateRef"`
	NetAmountTaxable    *float64 `json:"NetAmountTaxable,omitempty"`
	PercentBased        *bool    `json:"PercentBased,omitempty"`
	TaxInclusiveAmount  *float64 `json:"TaxInclusiveAmount,omitempty"`
	OverrideDeltaAmount *float64 `json:"OverrideDeltaAmount,omitempty"`
	TaxPercent          *float64 `json:"TaxPercent,omitempty"`
}

type MarkupInfo struct {
	PriceLevelRef          *Ref     `json:"PriceLevelRef,omitempty"`
	Percent                *float64 `json:"Percent,omitempty"`
	MarkUpIncomeAccountRef *Ref     `json:"MarkUpIncomeAccountRef,omitempty"`
}

type CustomField struct {
	DefinitionId string  `json:"DefinitionId"`
	StringValue  *string `json:"StringValue,omitempty"`
	Name         *string `json:"Name,omitempty"`
	Type         string  `json:"Type"`
}

type WebSiteAddress struct {
	URI *string `json:"URI,omitempty"`
}

type CreditCardPaymentForPayment struct {
	CreditChargeResponse *CreditChargeResponse `json:"CreditChargeResponse,omitempty"`
	CreditChargeInfo     *CreditChargeInfo     `json:"CreditChargeInfo,omitempty"`
}

type CreditChargeResponse struct {
	Status               *string    `json:"Status,omitempty"`   // CCPaymentStatusEnum: Completed, Unknown
	AuthCode             *string    `json:"AuthCode,omitempty"` // Max 100 chars
	TxnAuthorizationTime *time.Time `json:"TxnAuthorizationTime,omitempty"`
	CCTransId            *string    `json:"CCTransId,omitempty"` // Max 100 chars
}

type CreditChargeInfo struct {
	CcExpiryMonth  *int     `json:"CcExpiryMonth,omitempty"`
	ProcessPayment *bool    `json:"ProcessPayment,omitempty"`
	PostalCode     *string  `json:"PostalCode,omitempty"` // Max 30 chars
	Amount         *float64 `json:"Amount,omitempty"`
	NameOnAcct     *string  `json:"NameOnAcct,omitempty"`
	CcExpiryYear   *int     `json:"CcExpiryYear,omitempty"`
	Type           *string  `json:"Type,omitempty"`
	BillAddrStreet *string  `json:"BillAddrStreet,omitempty"` // Max 255 chars
}
