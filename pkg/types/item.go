package types

type Item struct {
	Id                   string                `json:"Id"`                         // Read-only, system defined, required for update
	ItemCategoryType     *string               `json:"ItemCategoryType,omitempty"` // Required, minorVersion: 3
	Name                 string                `json:"Name"`
	SyncToken            string                `json:"SyncToken"`                      // Required for update
	InvStartDate         string                `json:"InvStartDate"`                   // Conditionally required
	Type                 *string               `json:"Type,omitempty"`                 // Conditionally required, minorVersion: specified
	QtyOnHand            *float64              `json:"QtyOnHand,omitempty"`            // Conditionally required
	AssetAccountRef      *Ref                  `json:"AssetAccountRef,omitempty"`      // Conditionally required
	Sku                  *string               `json:"Sku,omitempty"`                  // Optional, minorVersion: 4
	SalesTaxIncluded     *bool                 `json:"SalesTaxIncluded,omitempty"`     // Optional
	TrackQtyOnHand       *bool                 `json:"TrackQtyOnHand,omitempty"`       // Optional
	SalesTaxCodeRef      *Ref                  `json:"SalesTaxCodeRef,omitempty"`      // Optional
	ClassRef             *Ref                  `json:"ClassRef,omitempty"`             // Optional, minorVersion: 41
	Source               *string               `json:"Source,omitempty"`               // Optional, minorVersion: 59
	PurchaseTaxIncluded  *bool                 `json:"PurchaseTaxIncluded,omitempty"`  // Optional
	Description          *string               `json:"Description,omitempty"`          // Optional
	AbatementRate        *float64              `json:"AbatementRate,omitempty"`        // Optional, minorVersion: 3
	SubItem              *bool                 `json:"SubItem,omitempty"`              // Optional
	Taxable              *bool                 `json:"Taxable,omitempty"`              // Optional
	UQCDisplayText       *string               `json:"UQCDisplayText,omitempty"`       // Optional, minorVersion: 33
	ReorderPoint         *float64              `json:"ReorderPoint,omitempty"`         // Optional
	PurchaseDesc         *string               `json:"PurchaseDesc,omitempty"`         // Optional
	MetaData             *ModificationMetaData `json:"MetaData,omitempty"`             // Optional
	PrefVendorRef        *Ref                  `json:"PrefVendorRef,omitempty"`        // Optional, minorVersion: 31
	Active               *bool                 `json:"Active,omitempty"`               // Optional
	UQCId                *string               `json:"UQCId,omitempty"`                // Optional, minorVersion: 33
	ReverseChargeRate    *float64              `json:"ReverseChargeRate,omitempty"`    // Optional, minorVersion: 3
	PurchaseTaxCodeRef   *Ref                  `json:"PurchaseTaxCodeRef,omitempty"`   // Optional
	ServiceType          *string               `json:"ServiceType,omitempty"`          // Optional, minorVersion: 3
	PurchaseCost         *float64              `json:"PurchaseCost,omitempty"`         // Optional
	ParentRef            *Ref                  `json:"ParentRef,omitempty"`            // Optional
	UnitPrice            *float64              `json:"UnitPrice,omitempty"`            // Optional
	FullyQualifiedName   string                `json:"FullyQualifiedName"`             // Read-only, system defined
	ExpenseAccountRef    *Ref                  `json:"ExpenseAccountRef,omitempty"`    // Optional for France locales
	Level                int                   `json:"Level"`                          // Read-only, system defined
	IncomeAccountRef     *Ref                  `json:"IncomeAccountRef,omitempty"`     // Conditionally required
	TaxClassificationRef *Ref                  `json:"TaxClassificationRef,omitempty"` // Optional, minorVersion: 34
}

type ItemPaginatedResponse struct {
	BasePaginatedResponse
	Item []Item `json:"Item"`
}

type ItemResponse struct {
	BaseResponse
	Item Item `json:"Item"`
}
