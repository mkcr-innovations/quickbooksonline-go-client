package types

type Account struct {
	BaseEntity
	Name                          string                `json:"Name"`
	AcctNum                       string                `json:"AcctNum"`
	CurrencyRef                   *Ref                  `json:"CurrencyRef,omitempty"`
	ParentRef                     *Ref                  `json:"ParentRef,omitempty"`
	Description                   *string               `json:"Description,omitempty"`
	Active                        *bool                 `json:"Active,omitempty"`
	MetaData                      *ModificationMetaData `json:"MetaData,omitempty"`
	SubAccount                    bool                  `json:"SubAccount"`
	Classification                string                `json:"Classification"`
	FullyQualifiedName            string                `json:"FullyQualifiedName"`
	TxnLocationType               *string               `json:"TxnLocationType,omitempty"`
	AccountType                   string                `json:"AccountType"`
	CurrentBalanceWithSubAccounts float64               `json:"CurrentBalanceWithSubAccounts"`
	AccountAlias                  *string               `json:"AccountAlias,omitempty"`
	TaxCodeRef                    *Ref                  `json:"TaxCodeRef,omitempty"`
	AccountSubType                string                `json:"AccountSubType"`
	CurrentBalance                float64               `json:"CurrentBalance"`
}

type AccountPaginatedResponse struct {
	BasePaginatedResponse
	Account []Account `json:"Account"`
}

type AccountResponse struct {
	BaseResponse
	Account Account `json:"Account"`
}
