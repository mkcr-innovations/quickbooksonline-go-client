package types

type Account struct {
	BaseEntity
	Name                          string                    `json:"Name"`
	AcctNum                       *string                   `json:"AcctNum,omitempty"`
	CurrencyRef                   *Ref                      `json:"CurrencyRef,omitempty"`
	ParentRef                     *Ref                      `json:"ParentRef,omitempty"`
	Description                   *string                   `json:"Description,omitempty"`
	Active                        *bool                     `json:"Active,omitempty"`
	MetaData                      *ModificationMetaData     `json:"MetaData,omitempty"`
	SubAccount                    bool                      `json:"SubAccount"`
	Classification                AccountClassificationEnum `json:"Classification"`
	FullyQualifiedName            string                    `json:"FullyQualifiedName"`
	TxnLocationType               *TxnLocationTypeEnum      `json:"TxnLocationType,omitempty"`
	AccountType                   AccountTypeEnum           `json:"AccountType"`
	CurrentBalanceWithSubAccounts float64                   `json:"CurrentBalanceWithSubAccounts"`
	AccountAlias                  *string                   `json:"AccountAlias,omitempty"`
	TaxCodeRef                    *Ref                      `json:"TaxCodeRef,omitempty"`
	AccountSubType                AccountSubTypeEnum        `json:"AccountSubType"`
	CurrentBalance                float64                   `json:"CurrentBalance"`
}

type AccountPaginatedResponse struct {
	BasePaginatedResponse
	Account []Account `json:"Account"`
}

type AccountResponse struct {
	BaseResponse
	Account Account `json:"Account"`
}

type TxnLocationTypeEnum string

const (
	WithinFrance        TxnLocationTypeEnum = "WithinFrance"
	FranceOverseas      TxnLocationTypeEnum = "FranceOverseas"
	OutsideFranceWithEU TxnLocationTypeEnum = "OutsideFranceWithEU"
	OutsideEU           TxnLocationTypeEnum = "OutsideEU"
)

type AccountClassificationEnum string

const (
	ASSET     AccountClassificationEnum = "Asset"
	EQUITY    AccountClassificationEnum = "Equity"
	EXPENSE   AccountClassificationEnum = "Expense"
	LIABILITY AccountClassificationEnum = "Liability"
	REVENUE   AccountClassificationEnum = "Revenue"
)

type AccountTypeEnum string

const (
	AccountTypeBank                  AccountTypeEnum = "Bank"
	AccountTypeOtherCurrentAsset     AccountTypeEnum = "Other Current Asset"
	AccountTypeFixedAsset            AccountTypeEnum = "Fixed Asset"
	AccountTypeOtherAsset            AccountTypeEnum = "Other Asset"
	AccountTypeAccountsReceivable    AccountTypeEnum = "Accounts Receivable"
	AccountTypeEquity                AccountTypeEnum = "Equity"
	AccountTypeExpense               AccountTypeEnum = "Expense"
	AccountTypeOtherExpense          AccountTypeEnum = "Other Expense"
	AccountTypeCostOfGoodsSold       AccountTypeEnum = "Cost of Goods Sold"
	AccountTypeAccountsPayable       AccountTypeEnum = "Accounts Payable"
	AccountTypeCreditCard            AccountTypeEnum = "Credit Card"
	AccountTypeLongTermLiability     AccountTypeEnum = "Long Term Liability"
	AccountTypeOtherCurrentLiability AccountTypeEnum = "Other Current Liability"
	AccountTypeIncome                AccountTypeEnum = "Income"
	AccountTypeOtherIncome           AccountTypeEnum = "Other Income"
)

type AccountSubTypeEnum string

const (
	// Bank SubTypes
	BankCashOnHand                 AccountSubTypeEnum = "CashOnHand"
	BankChecking                   AccountSubTypeEnum = "Checking"
	BankMoneyMarket                AccountSubTypeEnum = "MoneyMarket"
	BankRentsHeldInTrust           AccountSubTypeEnum = "RentsHeldInTrust"
	BankSavings                    AccountSubTypeEnum = "Savings"
	BankTrustAccounts              AccountSubTypeEnum = "TrustAccounts"
	BankCashAndCashEquivalents     AccountSubTypeEnum = "CashAndCashEquivalents"
	BankOtherEarMarkedBankAccounts AccountSubTypeEnum = "OtherEarMarkedBankAccounts"

	// OtherCurrentAsset SubTypes
	OtherCurrentAssetAllowanceForBadDebts                        AccountSubTypeEnum = "AllowanceForBadDebts"
	OtherCurrentAssetDevelopmentCosts                            AccountSubTypeEnum = "DevelopmentCosts"
	OtherCurrentAssetEmployeeCashAdvances                        AccountSubTypeEnum = "EmployeeCashAdvances"
	OtherCurrentAssetOtherCurrentAssets                          AccountSubTypeEnum = "OtherCurrentAssets"
	OtherCurrentAssetInventory                                   AccountSubTypeEnum = "Inventory"
	OtherCurrentAssetInvestment_MortgageRealEstateLoans          AccountSubTypeEnum = "Investment_MortgageRealEstateLoans"
	OtherCurrentAssetInvestment_Other                            AccountSubTypeEnum = "Investment_Other"
	OtherCurrentAssetInvestment_TaxExemptSecurities              AccountSubTypeEnum = "Investment_TaxExemptSecurities"
	OtherCurrentAssetInvestment_USGovernmentObligations          AccountSubTypeEnum = "Investment_USGovernmentObligations"
	OtherCurrentAssetLoansToOfficers                             AccountSubTypeEnum = "LoansToOfficers"
	OtherCurrentAssetLoansToOthers                               AccountSubTypeEnum = "LoansToOthers"
	OtherCurrentAssetLoansToStockholders                         AccountSubTypeEnum = "LoansToStockholders"
	OtherCurrentAssetPrepaidExpenses                             AccountSubTypeEnum = "PrepaidExpenses"
	OtherCurrentAssetRetainage                                   AccountSubTypeEnum = "Retainage"
	OtherCurrentAssetUndepositedFunds                            AccountSubTypeEnum = "UndepositedFunds"
	OtherCurrentAssetAssetsAvailableForSale                      AccountSubTypeEnum = "AssetsAvailableForSale"
	OtherCurrentAssetBalWithGovtAuthorities                      AccountSubTypeEnum = "BalWithGovtAuthorities"
	OtherCurrentAssetCalledUpShareCapitalNotPaid                 AccountSubTypeEnum = "CalledUpShareCapitalNotPaid"
	OtherCurrentAssetExpenditureAuthorisationsAndLettersOfCredit AccountSubTypeEnum = "ExpenditureAuthorisationsAndLettersOfCredit"
	OtherCurrentAssetGlobalTaxDeferred                           AccountSubTypeEnum = "GlobalTaxDeferred"
	OtherCurrentAssetGlobalTaxRefund                             AccountSubTypeEnum = "GlobalTaxRefund"
	OtherCurrentAssetInternalTransfers                           AccountSubTypeEnum = "InternalTransfers"
	OtherCurrentAssetOtherConsumables                            AccountSubTypeEnum = "OtherConsumables"
	OtherCurrentAssetProvisionsCurrentAssets                     AccountSubTypeEnum = "ProvisionsCurrentAssets"
	OtherCurrentAssetShortTermInvestmentsInRelatedParties        AccountSubTypeEnum = "ShortTermInvestmentsInRelatedParties"
	OtherCurrentAssetShortTermLoansAndAdvancesToRelatedParties   AccountSubTypeEnum = "ShortTermLoansAndAdvancesToRelatedParties"
	OtherCurrentAssetTradeAndOtherReceivables                    AccountSubTypeEnum = "TradeAndOtherReceivables"

	// FixedAsset SubTypes
	FixedAssetAccumulatedDepletion                     AccountSubTypeEnum = "AccumulatedDepletion"
	FixedAssetAccumulatedDepreciation                  AccountSubTypeEnum = "AccumulatedDepreciation"
	FixedAssetDepletableAssets                         AccountSubTypeEnum = "DepletableAssets"
	FixedAssetFixedAssetComputers                      AccountSubTypeEnum = "FixedAssetComputers"
	FixedAssetFixedAssetCopiers                        AccountSubTypeEnum = "FixedAssetCopiers"
	FixedAssetFixedAssetFurniture                      AccountSubTypeEnum = "FixedAssetFurniture"
	FixedAssetFixedAssetPhone                          AccountSubTypeEnum = "FixedAssetPhone"
	FixedAssetFixedAssetPhotoVideo                     AccountSubTypeEnum = "FixedAssetPhotoVideo"
	FixedAssetFixedAssetSoftware                       AccountSubTypeEnum = "FixedAssetSoftware"
	FixedAssetFixedAssetOtherToolsEquipment            AccountSubTypeEnum = "FixedAssetOtherToolsEquipment"
	FixedAssetFurnitureAndFixtures                     AccountSubTypeEnum = "FurnitureAndFixtures"
	FixedAssetLand                                     AccountSubTypeEnum = "Land"
	FixedAssetLeaseholdImprovements                    AccountSubTypeEnum = "LeaseholdImprovements"
	FixedAssetOtherFixedAssets                         AccountSubTypeEnum = "OtherFixedAssets"
	FixedAssetAccumulatedAmortization                  AccountSubTypeEnum = "AccumulatedAmortization"
	FixedAssetBuildings                                AccountSubTypeEnum = "Buildings"
	FixedAssetIntangibleAssets                         AccountSubTypeEnum = "IntangibleAssets"
	FixedAssetMachineryAndEquipment                    AccountSubTypeEnum = "MachineryAndEquipment"
	FixedAssetVehicles                                 AccountSubTypeEnum = "Vehicles"
	FixedAssetAssetsInCourseOfConstruction             AccountSubTypeEnum = "AssetsInCourseOfConstruction"
	FixedAssetCapitalWip                               AccountSubTypeEnum = "CapitalWip"
	FixedAssetCumulativeDepreciationOnIntangibleAssets AccountSubTypeEnum = "CumulativeDepreciationOnIntangibleAssets"
	FixedAssetIntangibleAssetsUnderDevelopment         AccountSubTypeEnum = "IntangibleAssetsUnderDevelopment"
	FixedAssetLandAsset                                AccountSubTypeEnum = "LandAsset"
	FixedAssetNonCurrentAssets                         AccountSubTypeEnum = "NonCurrentAssets"
	FixedAssetParticipatingInterests                   AccountSubTypeEnum = "ParticipatingInterests"
	FixedAssetProvisionsFixedAssets                    AccountSubTypeEnum = "ProvisionsFixedAssets"

	// OtherAsset SubTypes
	OtherAssetLeaseBuyout                              AccountSubTypeEnum = "LeaseBuyout"
	OtherAssetOtherLongTermAssets                      AccountSubTypeEnum = "OtherLongTermAssets"
	OtherAssetSecurityDeposits                         AccountSubTypeEnum = "SecurityDeposits"
	OtherAssetAccumulatedAmortizationOfOtherAssets     AccountSubTypeEnum = "AccumulatedAmortizationOfOtherAssets"
	OtherAssetGoodwill                                 AccountSubTypeEnum = "Goodwill"
	OtherAssetLicenses                                 AccountSubTypeEnum = "Licenses"
	OtherAssetOrganizationalCosts                      AccountSubTypeEnum = "OrganizationalCosts"
	OtherAssetAssetsHeldForSale                        AccountSubTypeEnum = "AssetsHeldForSale"
	OtherAssetAvailableForSaleFinancialAssets          AccountSubTypeEnum = "AvailableForSaleFinancialAssets"
	OtherAssetDeferredTax                              AccountSubTypeEnum = "DeferredTax"
	OtherAssetInvestments                              AccountSubTypeEnum = "Investments"
	OtherAssetLongTermInvestments                      AccountSubTypeEnum = "LongTermInvestments"
	OtherAssetLongTermLoansAndAdvancesToRelatedParties AccountSubTypeEnum = "LongTermLoansAndAdvancesToRelatedParties"
	OtherAssetOtherIntangibleAssets                    AccountSubTypeEnum = "OtherIntangibleAssets"
	OtherAssetOtherLongTermInvestments                 AccountSubTypeEnum = "OtherLongTermInvestments"
	OtherAssetOtherLongTermLoansAndAdvances            AccountSubTypeEnum = "OtherLongTermLoansAndAdvances"
	OtherAssetPrepaymentsAndAccruedIncome              AccountSubTypeEnum = "PrepaymentsAndAccruedIncome"
	OtherAssetProvisionsNonCurrentAssets               AccountSubTypeEnum = "ProvisionsNonCurrentAssets"

	// AccountsReceivable SubTypes
	AccountsReceivableAccountsReceivable AccountSubTypeEnum = "Accounts Receivable"

	// Equity SubTypes
	EquityOpeningBalanceEquity                  AccountSubTypeEnum = "OpeningBalanceEquity"
	EquityPartnersEquity                        AccountSubTypeEnum = "PartnersEquity"
	EquityRetainedEarnings                      AccountSubTypeEnum = "RetainedEarnings"
	EquityAccumulatedAdjustment                 AccountSubTypeEnum = "AccumulatedAdjustment"
	EquityOwnersEquity                          AccountSubTypeEnum = "OwnersEquity"
	EquityPaidInCapitalOrSurplus                AccountSubTypeEnum = "PaidInCapitalOrSurplus"
	EquityPartnerContributions                  AccountSubTypeEnum = "PartnerContributions"
	EquityPartnerDistributions                  AccountSubTypeEnum = "PartnerDistributions"
	EquityPreferredStock                        AccountSubTypeEnum = "PreferredStock"
	EquityCommonStock                           AccountSubTypeEnum = "CommonStock"
	EquityTreasuryStock                         AccountSubTypeEnum = "TreasuryStock"
	EquityEstimatedTaxes                        AccountSubTypeEnum = "EstimatedTaxes"
	EquityHealthcare                            AccountSubTypeEnum = "Healthcare"
	EquityPersonalIncome                        AccountSubTypeEnum = "PersonalIncome"
	EquityPersonalExpense                       AccountSubTypeEnum = "PersonalExpense"
	EquityAccumulatedOtherComprehensiveIncome   AccountSubTypeEnum = "AccumulatedOtherComprehensiveIncome"
	EquityCalledUpShareCapital                  AccountSubTypeEnum = "CalledUpShareCapital"
	EquityCapitalReserves                       AccountSubTypeEnum = "CapitalReserves"
	EquityDividendDisbursed                     AccountSubTypeEnum = "DividendDisbursed"
	EquityEquityInEarningsOfSubsiduaries        AccountSubTypeEnum = "EquityInEarningsOfSubsiduaries"
	EquityInvestmentGrants                      AccountSubTypeEnum = "InvestmentGrants"
	EquityMoneyReceivedAgainstShareWarrants     AccountSubTypeEnum = "MoneyReceivedAgainstShareWarrants"
	EquityOtherFreeReserves                     AccountSubTypeEnum = "OtherFreeReserves"
	EquityShareApplicationMoneyPendingAllotment AccountSubTypeEnum = "ShareApplicationMoneyPendingAllotment"
	EquityShareCapital                          AccountSubTypeEnum = "ShareCapital"
	EquityFunds                                 AccountSubTypeEnum = "Funds"

	// Expense SubTypes
	ExpenseAdvertisingPromotional                AccountSubTypeEnum = "AdvertisingPromotional"
	ExpenseBadDebts                              AccountSubTypeEnum = "BadDebts"
	ExpenseBankCharges                           AccountSubTypeEnum = "BankCharges"
	ExpenseCharitableContributions               AccountSubTypeEnum = "CharitableContributions"
	ExpenseCommissionsAndFees                    AccountSubTypeEnum = "CommissionsAndFees"
	ExpenseEntertainment                         AccountSubTypeEnum = "Entertainment"
	ExpenseEntertainmentMeals                    AccountSubTypeEnum = "EntertainmentMeals"
	ExpenseEquipmentRental                       AccountSubTypeEnum = "EquipmentRental"
	ExpenseFinanceCosts                          AccountSubTypeEnum = "FinanceCosts"
	ExpenseGlobalTaxExpense                      AccountSubTypeEnum = "GlobalTaxExpense"
	ExpenseInsurance                             AccountSubTypeEnum = "Insurance"
	ExpenseInterestPaid                          AccountSubTypeEnum = "InterestPaid"
	ExpenseLegalProfessionalFees                 AccountSubTypeEnum = "LegalProfessionalFees"
	ExpenseOfficeExpenses                        AccountSubTypeEnum = "OfficeExpenses"
	ExpenseOfficeGeneralAdministrativeExpenses   AccountSubTypeEnum = "OfficeGeneralAdministrativeExpenses"
	ExpenseOtherBusinessExpenses                 AccountSubTypeEnum = "OtherBusinessExpenses"
	ExpenseOtherMiscellaneousServiceCost         AccountSubTypeEnum = "OtherMiscellaneousServiceCost"
	ExpensePromotionalMeals                      AccountSubTypeEnum = "PromotionalMeals"
	ExpenseRentOrLeaseOfBuildings                AccountSubTypeEnum = "RentOrLeaseOfBuildings"
	ExpenseRepairMaintenance                     AccountSubTypeEnum = "RepairMaintenance"
	ExpenseShippingFreightDelivery               AccountSubTypeEnum = "ShippingFreightDelivery"
	ExpenseSuppliesMaterials                     AccountSubTypeEnum = "SuppliesMaterials"
	ExpenseTravel                                AccountSubTypeEnum = "Travel"
	ExpenseTravelMeals                           AccountSubTypeEnum = "TravelMeals"
	ExpenseAuto                                  AccountSubTypeEnum = "Auto"
	ExpenseCostOfLabor                           AccountSubTypeEnum = "CostOfLabor"
	ExpenseDuesSubscriptions                     AccountSubTypeEnum = "DuesSubscriptions"
	ExpensePayrollExpenses                       AccountSubTypeEnum = "PayrollExpenses"
	ExpenseTaxesPaid                             AccountSubTypeEnum = "TaxesPaid"
	ExpenseUnappliedCashBillPaymentExpense       AccountSubTypeEnum = "UnappliedCashBillPaymentExpense"
	ExpenseUtilities                             AccountSubTypeEnum = "Utilities"
	ExpenseAmortizationExpense                   AccountSubTypeEnum = "AmortizationExpense"
	ExpenseAppropriationsToDepreciation          AccountSubTypeEnum = "AppropriationsToDepreciation"
	ExpenseBorrowingCost                         AccountSubTypeEnum = "BorrowingCost"
	ExpenseDistributionCosts                     AccountSubTypeEnum = "DistributionCosts"
	ExpenseExternalServices                      AccountSubTypeEnum = "ExternalServices"
	ExpenseExtraordinaryCharges                  AccountSubTypeEnum = "ExtraordinaryCharges"
	ExpenseIncomeTaxExpense                      AccountSubTypeEnum = "IncomeTaxExpense"
	ExpenseLossOnDiscontinuedOperationsNetOfTax  AccountSubTypeEnum = "LossOnDiscontinuedOperationsNetOfTax"
	ExpenseManagementCompensation                AccountSubTypeEnum = "ManagementCompensation"
	ExpenseOtherCurrentOperatingCharges          AccountSubTypeEnum = "OtherCurrentOperatingCharges"
	ExpenseOtherExternalServices                 AccountSubTypeEnum = "OtherExternalServices"
	ExpenseOtherRentalCosts                      AccountSubTypeEnum = "OtherRentalCosts"
	ExpenseOtherSellingExpenses                  AccountSubTypeEnum = "OtherSellingExpenses"
	ExpenseProjectStudiesSurveysAssessments      AccountSubTypeEnum = "ProjectStudiesSurveysAssessments"
	ExpensePurchasesRebates                      AccountSubTypeEnum = "PurchasesRebates"
	ExpenseShippingAndDeliveryExpense            AccountSubTypeEnum = "ShippingAndDeliveryExpense"
	ExpenseStaffCosts                            AccountSubTypeEnum = "StaffCosts"
	ExpenseSundry                                AccountSubTypeEnum = "Sundry"
	ExpenseTravelExpensesGeneralAndAdminExpenses AccountSubTypeEnum = "TravelExpensesGeneralAndAdminExpenses"
	ExpenseTravelExpensesSellingExpense          AccountSubTypeEnum = "TravelExpensesSellingExpense"

	// OtherExpense SubTypes
	OtherExpensesDepreciation              AccountSubTypeEnum = "Depreciation"
	OtherExpensesExchangeGainOrLoss        AccountSubTypeEnum = "ExchangeGainOrLoss"
	OtherExpensesOtherMiscellaneousExpense AccountSubTypeEnum = "OtherMiscellaneousExpense"
	OtherExpensesPenaltiesSettlements      AccountSubTypeEnum = "PenaltiesSettlements"
	OtherExpensesAmortization              AccountSubTypeEnum = "Amortization"
	OtherExpensesGasAndFuel                AccountSubTypeEnum = "GasAndFuel"
	OtherExpensesHomeOffice                AccountSubTypeEnum = "HomeOffice"
	OtherExpensesHomeOwnerRentalInsurance  AccountSubTypeEnum = "HomeOwnerRentalInsurance"
	OtherExpensesOtherHomeOfficeExpenses   AccountSubTypeEnum = "OtherHomeOfficeExpenses"
	OtherExpensesMortgageInterest          AccountSubTypeEnum = "MortgageInterest"
	OtherExpensesRentAndLease              AccountSubTypeEnum = "RentAndLease"
	OtherExpensesRepairsAndMaintenance     AccountSubTypeEnum = "RepairsAndMaintenance"
	OtherExpensesParkingAndTolls           AccountSubTypeEnum = "ParkingAndTolls"
	OtherExpensesVehicle                   AccountSubTypeEnum = "Vehicle"
	OtherExpensesVehicleInsurance          AccountSubTypeEnum = "VehicleInsurance"
	OtherExpensesVehicleLease              AccountSubTypeEnum = "VehicleLease"
	OtherExpensesVehicleLoanInterest       AccountSubTypeEnum = "VehicleLoanInterest"
	OtherExpensesVehicleLoan               AccountSubTypeEnum = "VehicleLoan"
	OtherExpensesVehicleRegistration       AccountSubTypeEnum = "VehicleRegistration"
	OtherExpensesVehicleRepairs            AccountSubTypeEnum = "VehicleRepairs"
	OtherExpensesOtherVehicleExpenses      AccountSubTypeEnum = "OtherVehicleExpenses"
	OtherExpensesUtilities                 AccountSubTypeEnum = "Utilities"
	OtherExpensesWashAndRoadServices       AccountSubTypeEnum = "WashAndRoadServices"
	OtherExpensesDeferredTaxExpense        AccountSubTypeEnum = "DeferredTaxExpense"
	OtherExpensesDepletion                 AccountSubTypeEnum = "Depletion"
	OtherExpensesExceptionalItems          AccountSubTypeEnum = "ExceptionalItems"
	OtherExpensesExtraordinaryItems        AccountSubTypeEnum = "ExtraordinaryItems"
	OtherExpensesIncomeTaxOtherExpense     AccountSubTypeEnum = "IncomeTaxOtherExpense"
	OtherExpensesMatCredit                 AccountSubTypeEnum = "MatCredit"
	OtherExpensesPriorPeriodItems          AccountSubTypeEnum = "PriorPeriodItems"
	OtherExpensesTaxRoundoffGainOrLoss     AccountSubTypeEnum = "TaxRoundoffGainOrLoss"

	// CostOfGoodsSold SubTypes
	CostOfGoodsSoldEquipmentRentalCos         AccountSubTypeEnum = "EquipmentRentalCos"
	CostOfGoodsSoldOtherCostsOfServiceCos     AccountSubTypeEnum = "OtherCostsOfServiceCos"
	CostOfGoodsSoldShippingFreightDeliveryCos AccountSubTypeEnum = "ShippingFreightDeliveryCos"
	CostOfGoodsSoldSuppliesMaterialsCogs      AccountSubTypeEnum = "SuppliesMaterialsCogs"
	CostOfGoodsSoldCostOfLaborCos             AccountSubTypeEnum = "CostOfLaborCos"
	CostOfGoodsSoldCostOfSales                AccountSubTypeEnum = "CostOfSales"
	CostOfGoodsSoldFreightAndDeliveryCost     AccountSubTypeEnum = "FreightAndDeliveryCost"

	// AccountsPayable SubTypes
	AccountsPayableAccountsPayable                              AccountSubTypeEnum = "Accounts Payable"
	AccountsPayableOutstandingDuesMicroSmallEnterprise          AccountSubTypeEnum = "OutstandingDuesMicroSmallEnterprise"
	AccountsPayableOutstandingDuesOtherThanMicroSmallEnterprise AccountSubTypeEnum = "OutstandingDuesOtherThanMicroSmallEnterprise"

	// CreditCard SubTypes
	CreditCardCreditCard AccountSubTypeEnum = "Credit Card"

	// LongTermLiability SubTypes
	LongTermLiabilityNotesPayable                             AccountSubTypeEnum = "NotesPayable"
	LongTermLiabilityOtherLongTermLiabilities                 AccountSubTypeEnum = "OtherLongTermLiabilities"
	LongTermLiabilityShareholderNotesPayable                  AccountSubTypeEnum = "ShareholderNotesPayable"
	LongTermLiabilityAccrualsAndDeferredIncome                AccountSubTypeEnum = "AccrualsAndDeferredIncome"
	LongTermLiabilityAccruedLongLermLiabilities               AccountSubTypeEnum = "AccruedLongLermLiabilities"
	LongTermLiabilityAccruedVacationPayable                   AccountSubTypeEnum = "AccruedVacationPayable"
	LongTermLiabilityBankLoans                                AccountSubTypeEnum = "BankLoans"
	LongTermLiabilityDebtsRelatedToParticipatingInterests     AccountSubTypeEnum = "DebtsRelatedToParticipatingInterests"
	LongTermLiabilityDeferredTaxLiabilities                   AccountSubTypeEnum = "DeferredTaxLiabilities"
	LongTermLiabilityGovernmentAndOtherPublicAuthorities      AccountSubTypeEnum = "GovernmentAndOtherPublicAuthorities"
	LongTermLiabilityGroupAndAssociates                       AccountSubTypeEnum = "GroupAndAssociates"
	LongTermLiabilityLiabilitiesRelatedToAssetsHeldForSale    AccountSubTypeEnum = "LiabilitiesRelatedToAssetsHeldForSale"
	LongTermLiabilityLongTermBorrowings                       AccountSubTypeEnum = "LongTermBorrowings"
	LongTermLiabilityLongTermDebit                            AccountSubTypeEnum = "LongTermDebit"
	LongTermLiabilityLongTermEmployeeBenefitObligations       AccountSubTypeEnum = "LongTermEmployeeBenefitObligations"
	LongTermLiabilityObligationsUnderFinanceLeases            AccountSubTypeEnum = "ObligationsUnderFinanceLeases"
	LongTermLiabilityOtherLongTermProvisions                  AccountSubTypeEnum = "OtherLongTermProvisions"
	LongTermLiabilityProvisionForLiabilities                  AccountSubTypeEnum = "ProvisionForLiabilities"
	LongTermLiabilityProvisionsNonCurrentLiabilities          AccountSubTypeEnum = "ProvisionsNonCurrentLiabilities"
	LongTermLiabilityStaffAndRelatedLongTermLiabilityAccounts AccountSubTypeEnum = "StaffAndRelatedLongTermLiabilityAccounts"

	// OtherCurrentLiability SubTypes
	OtherCurrentLiabilityDirectDepositPayable                          AccountSubTypeEnum = "DirectDepositPayable"
	OtherCurrentLiabilityLineOfCredit                                  AccountSubTypeEnum = "LineOfCredit"
	OtherCurrentLiabilityLoanPayable                                   AccountSubTypeEnum = "LoanPayable"
	OtherCurrentLiabilityGlobalTaxPayable                              AccountSubTypeEnum = "GlobalTaxPayable"
	OtherCurrentLiabilityGlobalTaxSuspense                             AccountSubTypeEnum = "GlobalTaxSuspense"
	OtherCurrentLiabilityOtherCurrentLiabilities                       AccountSubTypeEnum = "OtherCurrentLiabilities"
	OtherCurrentLiabilityPayrollClearing                               AccountSubTypeEnum = "PayrollClearing"
	OtherCurrentLiabilityPayrollTaxPayable                             AccountSubTypeEnum = "PayrollTaxPayable"
	OtherCurrentLiabilityPrepaidExpensesPayable                        AccountSubTypeEnum = "PrepaidExpensesPayable"
	OtherCurrentLiabilityRentsInTrustLiability                         AccountSubTypeEnum = "RentsInTrustLiability"
	OtherCurrentLiabilityTrustAccountsLiabilities                      AccountSubTypeEnum = "TrustAccountsLiabilities"
	OtherCurrentLiabilityFederalIncomeTaxPayable                       AccountSubTypeEnum = "FederalIncomeTaxPayable"
	OtherCurrentLiabilityInsurancePayable                              AccountSubTypeEnum = "InsurancePayable"
	OtherCurrentLiabilitySalesTaxPayable                               AccountSubTypeEnum = "SalesTaxPayable"
	OtherCurrentLiabilityStateLocalIncomeTaxPayable                    AccountSubTypeEnum = "StateLocalIncomeTaxPayable"
	OtherCurrentLiabilityAccruedLiabilities                            AccountSubTypeEnum = "AccruedLiabilities"
	OtherCurrentLiabilityCurrentLiabilities                            AccountSubTypeEnum = "CurrentLiabilities"
	OtherCurrentLiabilityCurrentPortionEmployeeBenefitsObligations     AccountSubTypeEnum = "CurrentPortionEmployeeBenefitsObligations"
	OtherCurrentLiabilityCurrentPortionOfObligationsUnderFinanceLeases AccountSubTypeEnum = "CurrentPortionOfObligationsUnderFinanceLeases"
	OtherCurrentLiabilityCurrentTaxLiability                           AccountSubTypeEnum = "CurrentTaxLiability"
	OtherCurrentLiabilityDividendsPayable                              AccountSubTypeEnum = "DividendsPayable"
	OtherCurrentLiabilityDutiesAndTaxes                                AccountSubTypeEnum = "DutiesAndTaxes"
	OtherCurrentLiabilityInterestPayables                              AccountSubTypeEnum = "InterestPayables"
	OtherCurrentLiabilityProvisionForWarrantyObligations               AccountSubTypeEnum = "ProvisionForWarrantyObligations"
	OtherCurrentLiabilityProvisionsCurrentLiabilities                  AccountSubTypeEnum = "ProvisionsCurrentLiabilities"
	OtherCurrentLiabilityShortTermBorrowings                           AccountSubTypeEnum = "ShortTermBorrowings"
	OtherCurrentLiabilitySocialSecurityAgencies                        AccountSubTypeEnum = "SocialSecurityAgencies"
	OtherCurrentLiabilityStaffAndRelatedLiabilityAccounts              AccountSubTypeEnum = "StaffAndRelatedLiabilityAccounts"
	OtherCurrentLiabilitySundryDebtorsAndCreditors                     AccountSubTypeEnum = "SundryDebtorsAndCreditors"
	OtherCurrentLiabilityTradeAndOtherPayables                         AccountSubTypeEnum = "TradeAndOtherPayables"

	// Income SubTypes
	IncomeNonProfitIncome             AccountSubTypeEnum = "NonProfitIncome"
	IncomeOtherPrimaryIncome          AccountSubTypeEnum = "OtherPrimaryIncome"
	IncomeSalesOfProductIncome        AccountSubTypeEnum = "SalesOfProductIncome"
	IncomeServiceFeeIncome            AccountSubTypeEnum = "ServiceFeeIncome"
	IncomeDiscountsRefundsGiven       AccountSubTypeEnum = "DiscountsRefundsGiven"
	IncomeUnappliedCashPaymentIncome  AccountSubTypeEnum = "UnappliedCashPaymentIncome"
	IncomeCashReceiptIncome           AccountSubTypeEnum = "CashReceiptIncome"
	IncomeOperatingGrants             AccountSubTypeEnum = "OperatingGrants"
	IncomeOtherCurrentOperatingIncome AccountSubTypeEnum = "OtherCurrentOperatingIncome"
	IncomeOwnWorkCapitalized          AccountSubTypeEnum = "OwnWorkCapitalized"
	IncomeRevenueGeneral              AccountSubTypeEnum = "RevenueGeneral"
	IncomeSalesRetail                 AccountSubTypeEnum = "SalesRetail"
	IncomeSalesWholesale              AccountSubTypeEnum = "SalesWholesale"
	IncomeSavingsByTaxScheme          AccountSubTypeEnum = "SavingsByTaxScheme"

	// OtherIncome SubTypes
	OtherIncomeDividendIncome                     AccountSubTypeEnum = "DividendIncome"
	OtherIncomeInterestEarned                     AccountSubTypeEnum = "InterestEarned"
	OtherIncomeOtherMiscellaneousIncome           AccountSubTypeEnum = "OtherMiscellaneousIncome"
	OtherIncomeTaxExemptInterest                  AccountSubTypeEnum = "TaxExemptInterest"
	OtherIncomeGainLossOnSaleOfFixedAssets        AccountSubTypeEnum = "GainLossOnSaleOfFixedAssets"
	OtherIncomeGainLossOnSaleOfInvestments        AccountSubTypeEnum = "GainLossOnSaleOfInvestments"
	OtherIncomeLossOnDisposalOfAssets             AccountSubTypeEnum = "LossOnDisposalOfAssets"
	OtherIncomeOtherOperatingIncome               AccountSubTypeEnum = "OtherOperatingIncome"
	OtherIncomeUnrealisedLossOnSecuritiesNetOfTax AccountSubTypeEnum = "UnrealisedLossOnSecuritiesNetOfTax"
)
