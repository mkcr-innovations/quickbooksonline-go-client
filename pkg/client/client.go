package client

import (
	"fmt"

	"github.com/mkcr-innovations/quickbooksonline-go-client/internal/services"
	"github.com/mkcr-innovations/quickbooksonline-go-client/pkg/types"
)

type Client struct {
	baseURL    string
	realmId    string
	httpClient types.HttpClientInterface

	// Services
	Account           *services.BaseService[types.AccountResponse, types.AccountPaginatedResponse]
	Attachable        *services.BaseService[types.AttachableResponse, types.AttachablePaginatedResponse]
	Bill              *services.BaseService[types.BillResponse, types.BillPaginatedResponse]
	BillPayment       *services.BaseService[types.BillPaymentResponse, types.BillPaymentPaginatedResponse]
	Budget            *services.BaseService[types.BudgetResponse, types.BudgetPaginatedResponse]
	Class             *services.BaseService[types.ClassResponse, types.ClassPaginatedResponse]
	CompanyInfo       *services.BaseService[types.CompanyInfoResponse, types.CompanyInfoPaginatedResponse]
	CreditCardPayment *services.BaseService[types.CreditCardPaymentResponse, types.CreditCardPaymentPaginatedResponse]
	CreditMemo        *services.BaseService[types.CreditMemoResponse, types.CreditMemoPaginatedResponse]
	Customer          *services.BaseService[types.CustomerResponse, types.CustomerPaginatedResponse]
	CustomerType      *services.BaseService[types.CustomerTypeResponse, types.CustomerTypePaginatedResponse]
	Department        *services.BaseService[types.DepartmentResponse, types.DepartmentPaginatedResponse]
	Deposit           *services.BaseService[types.DepositResponse, types.DepositPaginatedResponse]
	Employee          *services.BaseService[types.EmployeeResponse, types.EmployeePaginatedResponse]
	Estimate          *services.BaseService[types.EstimateResponse, types.EstimatePaginatedResponse]
	Invoice           *services.BaseService[types.InvoiceResponse, types.InvoicePaginatedResponse]
	Item              *services.BaseService[types.ItemResponse, types.ItemPaginatedResponse]
	JournalCode       *services.BaseService[types.JournalCodeResponse, types.JournalCodePaginatedResponse]
	JournalEntry      *services.BaseService[types.JournalEntryResponse, types.JournalEntryPaginatedResponse]
	Payment           *services.BaseService[types.PaymentResponse, types.PaymentPaginatedResponse]
	Purchase          *services.BaseService[types.PurchaseResponse, types.PurchasePaginatedResponse]
	PurchaseOrder     *services.BaseService[types.PurchaseOrderResponse, types.PurchaseOrderPaginatedResponse]
	RefundReceipt     *services.BaseService[types.RefundReceiptResponse, types.RefundReceiptPaginatedResponse]
	ReimburseCharge   *services.BaseService[types.ReimburseChargeResponse, types.ReimburseChargePaginatedResponse]
	SalesReceipt      *services.BaseService[types.SalesReceiptResponse, types.SalesReceiptPaginatedResponse]
	TaxRate           *services.BaseService[types.TaxRateResponse, types.TaxRatePaginatedResponse]
	Transfer          *services.BaseService[types.TransferResponse, types.TransferPaginatedResponse]
	Vendor            *services.BaseService[types.VendorResponse, types.VendorPaginatedResponse]
	VendorCredit      *services.BaseService[types.VendorCreditResponse, types.VendorCreditPaginatedResponse]

	// Report Services
	ProfitAndLoss *services.ReportService
	BalanceSheet  *services.ReportService
	CashFlow      *services.ReportService
	GeneralLedger *services.ReportService
}

// baseEndpoint is the endpoint for the realm / account
func (c *Client) baseEndpoint() string {
	return fmt.Sprintf("%s/v3/company/%s", c.baseURL, c.realmId)
}

type Option func(*Client)

// WithSandbox overrides the baseURL to that of Sandbox
func WithSandbox() Option {
	return func(c *Client) {
		c.baseURL = "https://sandbox-quickbooks.api.intuit.com"
	}
}

// WithCustomBaseURL allows overriding the default endpoint
func WithCustomBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// NewClient - Use golang.org/x/oauth2 to generate the httpClient and pass it here
func NewClient(httpClient types.HttpClientInterface, realmId string, opts ...Option) *Client {
	client := Client{
		baseURL:    "https://quickbooks.api.intuit.com",
		realmId:    realmId,
		httpClient: httpClient,
	}

	// Apply any provided options to override defaults
	for _, opt := range opts {
		opt(&client)
	}

	// Register services
	client.Account = services.NewBaseService[types.AccountResponse, types.AccountPaginatedResponse](client.httpClient, client.baseEndpoint(), "account")
	client.Attachable = services.NewBaseService[types.AttachableResponse, types.AttachablePaginatedResponse](client.httpClient, client.baseEndpoint(), "attachable")
	client.Bill = services.NewBaseService[types.BillResponse, types.BillPaginatedResponse](client.httpClient, client.baseEndpoint(), "bill")
	client.BillPayment = services.NewBaseService[types.BillPaymentResponse, types.BillPaymentPaginatedResponse](client.httpClient, client.baseEndpoint(), "billpayment")
	client.Budget = services.NewBaseService[types.BudgetResponse, types.BudgetPaginatedResponse](client.httpClient, client.baseEndpoint(), "budget")
	client.Class = services.NewBaseService[types.ClassResponse, types.ClassPaginatedResponse](client.httpClient, client.baseEndpoint(), "class")
	client.CompanyInfo = services.NewBaseService[types.CompanyInfoResponse, types.CompanyInfoPaginatedResponse](client.httpClient, client.baseEndpoint(), "companyinfo")
	client.CreditCardPayment = services.NewBaseService[types.CreditCardPaymentResponse, types.CreditCardPaymentPaginatedResponse](client.httpClient, client.baseEndpoint(), "creditcardpayment")
	client.CreditMemo = services.NewBaseService[types.CreditMemoResponse, types.CreditMemoPaginatedResponse](client.httpClient, client.baseEndpoint(), "creditmemo")
	client.Customer = services.NewBaseService[types.CustomerResponse, types.CustomerPaginatedResponse](client.httpClient, client.baseEndpoint(), "customer")
	client.CustomerType = services.NewBaseService[types.CustomerTypeResponse, types.CustomerTypePaginatedResponse](client.httpClient, client.baseEndpoint(), "customertype")
	client.Department = services.NewBaseService[types.DepartmentResponse, types.DepartmentPaginatedResponse](client.httpClient, client.baseEndpoint(), "department")
	client.Deposit = services.NewBaseService[types.DepositResponse, types.DepositPaginatedResponse](client.httpClient, client.baseEndpoint(), "deposit")
	client.Employee = services.NewBaseService[types.EmployeeResponse, types.EmployeePaginatedResponse](client.httpClient, client.baseEndpoint(), "employee")
	client.Estimate = services.NewBaseService[types.EstimateResponse, types.EstimatePaginatedResponse](client.httpClient, client.baseEndpoint(), "estimate")
	client.Invoice = services.NewBaseService[types.InvoiceResponse, types.InvoicePaginatedResponse](client.httpClient, client.baseEndpoint(), "invoice")
	client.Item = services.NewBaseService[types.ItemResponse, types.ItemPaginatedResponse](client.httpClient, client.baseEndpoint(), "item")
	client.JournalCode = services.NewBaseService[types.JournalCodeResponse, types.JournalCodePaginatedResponse](client.httpClient, client.baseEndpoint(), "journalcode")
	client.JournalEntry = services.NewBaseService[types.JournalEntryResponse, types.JournalEntryPaginatedResponse](client.httpClient, client.baseEndpoint(), "journalentry")
	client.Payment = services.NewBaseService[types.PaymentResponse, types.PaymentPaginatedResponse](client.httpClient, client.baseEndpoint(), "payment")
	client.Purchase = services.NewBaseService[types.PurchaseResponse, types.PurchasePaginatedResponse](client.httpClient, client.baseEndpoint(), "purchase")
	client.PurchaseOrder = services.NewBaseService[types.PurchaseOrderResponse, types.PurchaseOrderPaginatedResponse](client.httpClient, client.baseEndpoint(), "purchaseorder")
	client.RefundReceipt = services.NewBaseService[types.RefundReceiptResponse, types.RefundReceiptPaginatedResponse](client.httpClient, client.baseEndpoint(), "refundreceipt")
	client.ReimburseCharge = services.NewBaseService[types.ReimburseChargeResponse, types.ReimburseChargePaginatedResponse](client.httpClient, client.baseEndpoint(), "reimbursecharge")
	client.SalesReceipt = services.NewBaseService[types.SalesReceiptResponse, types.SalesReceiptPaginatedResponse](client.httpClient, client.baseEndpoint(), "salesreceipt")
	client.TaxRate = services.NewBaseService[types.TaxRateResponse, types.TaxRatePaginatedResponse](client.httpClient, client.baseEndpoint(), "taxrate")
	client.Transfer = services.NewBaseService[types.TransferResponse, types.TransferPaginatedResponse](client.httpClient, client.baseEndpoint(), "transfer")
	client.Vendor = services.NewBaseService[types.VendorResponse, types.VendorPaginatedResponse](client.httpClient, client.baseEndpoint(), "vendor")
	client.VendorCredit = services.NewBaseService[types.VendorCreditResponse, types.VendorCreditPaginatedResponse](client.httpClient, client.baseEndpoint(), "vendorcredit")

	// Register Report services
	client.ProfitAndLoss = services.NewReportService(client.httpClient, client.baseEndpoint(), "ProfitAndLoss")
	client.BalanceSheet = services.NewReportService(client.httpClient, client.baseEndpoint(), "BalanceSheet")
	client.CashFlow = services.NewReportService(client.httpClient, client.baseEndpoint(), "CashFlow")
	client.GeneralLedger = services.NewReportService(client.httpClient, client.baseEndpoint(), "GeneralLedger")
	return &client
}
