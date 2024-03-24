package services

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/mkcr-innovations/quickbooksonline-go-client/pkg/types"
)

// ReportService is a service type shared across all types of reports
type ReportService struct {
	httpClient   types.HttpClientInterface
	baseEndpoint string
	entity       string
}

// NewReportService creates a new instance of ReportService with the specified entity.
func NewReportService(httpClient types.HttpClientInterface, baseEndpoint, entity string) *ReportService {
	return &ReportService{
		httpClient:   httpClient,
		baseEndpoint: baseEndpoint,
		entity:       entity,
	}
}

// Read directly fetches an entity by ID from the QuickBooks API.
func (s *ReportService) Query(queryParams map[string][]string) (*types.Report, error) {
	// Parse the URL from the string
	u, err := url.Parse(fmt.Sprintf("%s/reports/%s", s.baseEndpoint, s.entity))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for key, values := range queryParams {
		for _, value := range values {
			q.Add(key, value) // Use Add to support multiple values for the same key
		}
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response types.Report
	err = HandleResponse(resp, &response)
	return &response, err
}
