package services

import (
	"fmt"
	"net/http"

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

// Entity returns the entity name associated with the service.
func (s *ReportService) Entity() string {
	return s.entity
}

// Read directly fetches an entity by ID from the QuickBooks API.
func (s *ReportService) Query() (*types.Report, error) {
	url := fmt.Sprintf("%s/reports/%s", s.baseEndpoint, s.entity)
	req, err := http.NewRequest("GET", url, nil)
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
