package types

import (
	"net/http"
)

// HttpClientInterface abstracts the HTTP client
type HttpClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
