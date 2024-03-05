package types

type Line struct {
	Id          string       `json:"Id"`
	DetailType  *string      `json:"DetailType,omitempty"` // Required except for Payment object
	Amount      *float64     `json:"Amount,omitempty"`
	LinkedTxn   []LinkedTxn  `json:"LinkedTxn,omitempty"`
	Description *string      `json:"Description,omitempty"`
	LineNum     *int         `json:"LineNum,omitempty"`
	ProjectRef  *Ref         `json:"ProjectRef,omitempty"` // Optional, minorVersion: 69
	CustomField *CustomField `json:"CustomField,omitempty"`

	AccountBasedExpenseLineDetail *AccountBasedExpenseLineDetail `json:"AccountBasedExpenseLineDetail"`
	ItemBasedExpenseLineDetail    *ItemBasedExpenseLineDetail    `json:"ItemBasedExpenseLineDetail"`
	SalesItemLineDetail           *SalesItemLineDetail           `json:"SalesItemLineDetail,omitempty"`
	GroupLineDetail               *GroupLineDetail               `json:"GroupLineDetail,omitempty"`
	DescriptionLineDetail         *DescriptionLineDetail         `json:"DescriptionLineDetail,omitempty"`
	DiscountLineDetail            *DiscountLineDetail            `json:"DiscountLineDetail,omitempty"`
	SubtotalLineDetail            *SubtotalLineDetail            `json:"SubtotalLineDetail,omitempty"`
	JournalEntryLineDetail        *JournalEntryLineDetail        `json:"JournalEntryLineDetail,omitempty"`
	DepositLineDetail             *DepositLineDetail             `json:"DepositLineDetail,omitempty"`
	ReimburseLineDetail           *ReimburseLineDetail           `json:"ReimburseLineDetail,omitempty"`
}

type ReimburseLineDetail struct {
	ClassRef           *Ref     `json:"ClassRef,omitempty"`           // Optional
	TaxCodeRef         *Ref     `json:"TaxCodeRef,omitempty"`         // Optional
	DiscountAccountRef *Ref     `json:"DiscountAccountRef,omitempty"` // Optional
	PercentBased       *bool    `json:"PercentBased,omitempty"`       // Optional
	DiscountPercent    *float64 `json:"DiscountPercent,omitempty"`    // Optional
}

type DepositLineDetail struct {
	AccountRef       Ref     `json:"AccountRef"`
	PaymentMethodRef *Ref    `json:"PaymentMethodRef,omitempty"`
	ClassRef         *Ref    `json:"ClassRef,omitempty"`
	CheckNum         *string `json:"CheckNum,omitempty"`
	TaxCodeRef       *Ref    `json:"TaxCodeRef,omitempty"`
	TaxApplicableOn  *string `json:"TaxApplicableOn,omitempty"`
	TxnType          *string `json:"TxnType,omitempty"`
	Entity           *Ref    `json:"Entity,omitempty"`
}

type JournalEntryLineDetail struct {
	JournalCodeRef  *Ref     `json:"JournalCodeRef,omitempty"`  // Required for France locales, minorVersion: 5
	PostingType     string   `json:"PostingType"`               // Required
	AccountRef      Ref      `json:"AccountRef"`                // Required
	TaxApplicableOn *string  `json:"TaxApplicableOn,omitempty"` // Conditionally required
	Entity          *Entity  `json:"Entity,omitempty"`          // Conditionally required
	TaxInclusiveAmt *float64 `json:"TaxInclusiveAmt,omitempty"` // Optional, minorVersion: 53
	ClassRef        *Ref     `json:"ClassRef,omitempty"`        // Optional
	DepartmentRef   *Ref     `json:"DepartmentRef,omitempty"`   // Optional
	TaxCodeRef      *Ref     `json:"TaxCodeRef,omitempty"`      // Optional
	BillableStatus  *string  `json:"BillableStatus,omitempty"`  // Optional
	TaxAmount       *float64 `json:"TaxAmount,omitempty"`       // Optional
}

type Entity struct {
	EntityRef Ref    `json:"EntityRef"`
	Type      string `json:"Type"`
}

type SalesItemLineDetail struct {
	TaxInclusiveAmt      *float64    `json:"TaxInclusiveAmt,omitempty"`
	DiscountAmt          *float64    `json:"DiscountAmt,omitempty"`
	ItemRef              *Ref        `json:"ItemRef,omitempty"`
	ClassRef             *Ref        `json:"ClassRef,omitempty"`
	TaxCodeRef           *Ref        `json:"TaxCodeRef,omitempty"`
	MarkupInfo           *MarkupInfo `json:"MarkupInfo,omitempty"`
	ItemAccountRef       *Ref        `json:"ItemAccountRef,omitempty"`
	ServiceDate          *string     `json:"ServiceDate,omitempty"`
	DiscountRate         *float64    `json:"DiscountRate,omitempty"`
	Qty                  *float64    `json:"Qty,omitempty"`
	UnitPrice            *float64    `json:"UnitPrice,omitempty"`
	TaxClassificationRef *Ref        `json:"TaxClassificationRef,omitempty"`
}

type AccountBasedExpenseLineDetail struct {
	AccountRef      Ref         `json:"AccountRef"`
	TaxAmount       *float64    `json:"TaxAmount,omitempty"`
	TaxInclusiveAmt *float64    `json:"TaxInclusiveAmt,omitempty"`
	ClassRef        *Ref        `json:"ClassRef,omitempty"`
	TaxCodeRef      *Ref        `json:"TaxCodeRef,omitempty"`
	MarkupInfo      *MarkupInfo `json:"MarkupInfo,omitempty"`
	BillableStatus  *string     `json:"BillableStatus,omitempty"`
	CustomerRef     *Ref        `json:"CustomerRef,omitempty"`
}

type ItemBasedExpenseLineDetail struct {
	TaxInclusiveAmt *float64    `json:"TaxInclusiveAmt,omitempty"`
	ItemRef         *Ref        `json:"ItemRef,omitempty"`
	CustomerRef     *Ref        `json:"CustomerRef,omitempty"`
	PriceLevelRef   *Ref        `json:"PriceLevelRef,omitempty"`
	ClassRef        *Ref        `json:"ClassRef,omitempty"`
	TaxCodeRef      *Ref        `json:"TaxCodeRef,omitempty"`
	MarkupInfo      *MarkupInfo `json:"MarkupInfo,omitempty"`
	BillableStatus  *string     `json:"BillableStatus,omitempty"`
	Qty             *float64    `json:"Qty,omitempty"`
	UnitPrice       *float64    `json:"UnitPrice,omitempty"`
}

type GroupLineDetail struct {
	Quantity     *float64 `json:"Quantity,omitempty"`
	Lines        []Line   `json:"Lines,omitempty"`
	GroupItemRef *Ref     `json:"GroupItemRef,omitempty"`
}

type DescriptionLineDetail struct {
	TaxCodeRef  *Ref    `json:"TaxCodeRef,omitempty"`
	ServiceDate *string `json:"ServiceDate,omitempty"`
}

type DiscountLineDetail struct {
	ClassRef           *Ref     `json:"ClassRef,omitempty"`
	TaxCodeRef         *Ref     `json:"TaxCodeRef,omitempty"`
	DiscountAccountRef *Ref     `json:"DiscountAccountRef,omitempty"`
	PercentBased       *bool    `json:"PercentBased,omitempty"`
	DiscountPercent    *float64 `json:"DiscountPercent,omitempty"`
}

type SubtotalLineDetail struct {
	ItemRef *Ref `json:"ItemRef,omitempty"`
}
