package services

import (
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

// Entity returns the entity name associated with the service.
func (s *BaseService[Response, PaginatedResponse]) Entity() string {
	return s.entity
}

// Query initializes the query with the target entity.
func (s *BaseService[Response, PaginatedResponse]) Query() *QueryBuilder[PaginatedResponse] {
	return NewQueryBuilder[PaginatedResponse](s.httpClient, s.baseEndpoint, s.Entity())
}

// Read directly fetches an entity by ID from the QuickBooks API.
func (s *BaseService[Response, PaginatedResponse]) Read(id string) (*Response, error) {
	return ReadHandler[Response](id, s.Entity(), s.httpClient, s.baseEndpoint)
}
