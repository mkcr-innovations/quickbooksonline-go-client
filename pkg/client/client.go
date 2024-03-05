package client

import (
	"fmt"
	"quickbooksonline-go-client/internal/services"
	"quickbooksonline-go-client/pkg/types"
)

type Client struct {
	baseURL    string
	realmId    string
	httpClient types.HttpClientInterface

	// Cache services
	Account    *services.BaseService[types.AccountResponse, types.AccountPaginatedResponse]
	Attachable *services.BaseService[types.AttachableResponse, types.AttachablePaginatedResponse]
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
	return &client
}
