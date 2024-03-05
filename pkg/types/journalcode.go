package types

type JournalCode struct {
	BaseEntity
	Name        string                `json:"Name"`
	Description *string               `json:"Description,omitempty"`
	CustomField []CustomField         `json:"CustomField,omitempty"`
	Type        *JournalCodeTypeEnum  `json:"Type,omitempty"`
	MetaData    *ModificationMetaData `json:"MetaData,omitempty"`
}

type JournalCodePaginatedResponse struct {
	BasePaginatedResponse
	JournalCode []JournalCode `json:"JournalCode"`
}

type JournalCodeResponse struct {
	BaseResponse
	JournalCode JournalCode `json:"JournalCode"`
}

type JournalCodeTypeEnum string

const (
	ExpensesJournalCodeType JournalCodeTypeEnum = "Expenses"
	SalesJournalCodeType    JournalCodeTypeEnum = "Sales"
	BankJournalCodeType     JournalCodeTypeEnum = "Bank"
	NouveauxJournalCodeType JournalCodeTypeEnum = "Nouveaux"
	WagesJournalCodeType    JournalCodeTypeEnum = "Wages"
	CashJournalCodeType     JournalCodeTypeEnum = "Cash"
	OthersJournalCodeType   JournalCodeTypeEnum = "Others"
)
