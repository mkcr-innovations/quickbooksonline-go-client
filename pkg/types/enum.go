package types

type GlobalTaxCalculationEnum string

const (
	GlobalTaxCalculationTaxExcluded   GlobalTaxCalculationEnum = "TaxExcluded"
	GlobalTaxCalculationTaxInclusive  GlobalTaxCalculationEnum = "TaxInclusive"
	GlobalTaxCalculationNotApplicable GlobalTaxCalculationEnum = "NotApplicable"
)

type TransactionLocationTypeEnum string

const (
	TransactionLocationTypeWithinFrance        TransactionLocationTypeEnum = "WithinFrance"
	TransactionLocationTypeFranceOverseas      TransactionLocationTypeEnum = "FranceOverseas"
	TransactionLocationTypeOutsideFranceWithEU TransactionLocationTypeEnum = "OutsideFranceWithEU"
	TransactionLocationTypeOutsideEU           TransactionLocationTypeEnum = "OutsideEU"
)

type PrintStatusEnum string

const (
	NotSetPrintStatus        PrintStatusEnum = "NotSet"
	NeedToPrintPrintStatus   PrintStatusEnum = "NeedToPrint"
	PrintCompletePrintStatus PrintStatusEnum = "PrintComplete"
)

type EmailStatusEnum string

const (
	NotSetEmailStatus     EmailStatusEnum = "NotSet"
	NeedToSendEmailStatus EmailStatusEnum = "NeedToSend"
	EmailSentEmailStatus  EmailStatusEnum = "EmailSent"
)

type SourceEnum string

const QBCommerceCustomerSource SourceEnum = "QBCommerce"

type PaymentTypeEnum string

const (
	CashPaymentPayType       PaymentTypeEnum = "Cash"
	CheckPaymentPayType      PaymentTypeEnum = "Check"
	CreditCardPaymentPayType PaymentTypeEnum = "CreditCard"
	OtherPaymentPayType      PaymentTypeEnum = "Other"
)
