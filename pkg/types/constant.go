package types

import "golang.org/x/oauth2"

type Scope string

const (
	AccountingScope Scope = "com.intuit.quickbooks.accounting"
	PaymentScope    Scope = "com.intuit.quickbooks.payment"
	OpenIdScope     Scope = "openid"
)

var OAuth2Endpoint = oauth2.Endpoint{
	AuthURL:  "https://appcenter.intuit.com/connect/oauth2",
	TokenURL: "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer",
}
