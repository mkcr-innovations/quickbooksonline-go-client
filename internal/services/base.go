package services

import (
	"fmt"
	"net/http"

	"github.com/mkcr-innovations/quickbooksonline-go-client/pkg/types"
)

// BaseService is a generic service type with shared functionality
type BaseService[Response, PaginatedResponse any] struct {
	httpClient   types.HttpClientInterface
	baseEndpoint string
	entity       string
}

// NewBaseService creates a new instance of BaseService with the specified entity.
func NewBaseService[Response, PaginatedResponse any](httpClient types.HttpClientInterface, baseEndpoint, entity string) *BaseService[Response, PaginatedResponse] {
	return &BaseService[Response, PaginatedResponse]{
		httpClient:   httpClient,
		baseEndpoint: baseEndpoint,
		entity:       entity,
	}
}

// Query initializes the query with the target entity.
func (s *BaseService[Response, PaginatedResponse]) Query() *QueryBuilder[PaginatedResponse] {
	return NewQueryBuilder[PaginatedResponse](s.httpClient, s.baseEndpoint, s.entity)
}

// Read directly fetches an entity by ID from the QuickBooks API.
func (s *BaseService[Response, PaginatedResponse]) Read(id string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/%s", s.baseEndpoint, s.entity, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = HandleResponse(resp, &response)
	return &response, err
}
