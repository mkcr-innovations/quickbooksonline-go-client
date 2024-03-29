package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/mkcr-innovations/quickbooksonline-go-client/pkg/types"
)

// QueryBuilder struct to hold parts of our SQL query
type QueryBuilder[R any] struct {
	httpClient   types.HttpClientInterface
	baseEndpoint string
	entity       string

	count          bool
	selectFields   []string
	whereClauses   []string
	startPosition  int
	maxResults     int
	orderByClauses []string
}

func NewQueryBuilder[R any](httpClient types.HttpClientInterface, baseEndpoint string, entity string) *QueryBuilder[R] {
	return &QueryBuilder[R]{
		httpClient:   httpClient,
		baseEndpoint: baseEndpoint,
		entity:       entity,
	}
}

// Select specifies the fields to select
func (qb *QueryBuilder[R]) Select(fields ...string) *QueryBuilder[R] {
	qb.selectFields = fields
	return qb
}

// Where adds a WHERE clause using a condition
func (qb *QueryBuilder[R]) Where(condition string) *QueryBuilder[R] {
	qb.whereClauses = append(qb.whereClauses, condition)
	return qb
}

// StartPosition specifies the STARTPOSITION value
func (qb *QueryBuilder[R]) StartPosition(position int) *QueryBuilder[R] {
	qb.startPosition = position
	return qb
}

// MaxResults specifies the MAXRESULTS value
func (qb *QueryBuilder[R]) MaxResults(max int) *QueryBuilder[R] {
	qb.maxResults = max
	return qb
}

// OrderBy adds ORDER BY clause
func (qb *QueryBuilder[R]) OrderBy(field string, direction types.OrderByDirection) *QueryBuilder[R] {
	qb.orderByClauses = append(qb.orderByClauses, fmt.Sprintf("%s %s", field, direction))
	return qb
}

// Build assembles the SQL query
func (qb *QueryBuilder[R]) build() string {
	var query = "SELECT"
	if qb.count {
		query = fmt.Sprintf("%s COUNT(*)", query)
	} else {
		if len(qb.selectFields) == 0 {
			query = fmt.Sprintf("%s *", query)
		} else {
			query = fmt.Sprintf("%s %s", query, strings.Join(qb.selectFields, ", "))
		}
	}

	query = fmt.Sprintf("%s FROM %s", query, qb.entity)

	if len(qb.whereClauses) > 0 {
		query += " WHERE " + strings.Join(qb.whereClauses, " AND ")
	}
	if len(qb.orderByClauses) > 0 {
		query += " ORDER BY " + strings.Join(qb.orderByClauses, ", ")
	}
	if qb.startPosition > 0 {
		query += fmt.Sprintf(" STARTPOSITION %d", qb.startPosition)
	}
	if qb.maxResults > 0 {
		query += fmt.Sprintf(" MAXRESULTS %d", qb.maxResults)
	}
	return query
}

func (qb *QueryBuilder[R]) Exec() (*types.QueryResponse[R], error) {

	data := url.Values{}
	data.Set("query", qb.build())

	url := fmt.Sprintf("%s/query?%s", qb.baseEndpoint, data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := qb.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	response := types.QueryResponse[R]{}
	err = HandleResponse(resp, &response)
	return &response, err
}

// Count
func (qb *QueryBuilder[R]) ExecCount() (*types.QueryCountResponse, error) {
	qb.count = true

	data := url.Values{}
	data.Set("query", qb.build())

	url := fmt.Sprintf("%s/query?%s", qb.baseEndpoint, data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := qb.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	response := types.QueryCountResponse{}
	err = HandleResponse(resp, &response)
	return &response, err
}
