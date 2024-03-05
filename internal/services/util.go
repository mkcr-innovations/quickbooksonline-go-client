package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/mkcr-innovations/quickbooksonline-go-client/pkg/types"
)

func HandleResponse(resp *http.Response, respVar interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, &respVar); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return nil
}

func ReadHandler[T any](id string, entity string, httpClient types.HttpClientInterface, baseEndpoint string) (*T, error) {
	url := fmt.Sprintf("%s/%s/%s", baseEndpoint, entity, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response T
	err = HandleResponse(resp, &response)
	return &response, err
}

// NewService creates a new service of any type that includes an httpClient and baseEndpoint.
// T must be a pointer to a struct type that has `httpClient` and `baseEndpoint` fields.
func NewService[T any](httpClient types.HttpClientInterface, baseEndpoint string) *T {
	t := new(T) // Create a new instance of the type T
	v := reflect.ValueOf(t).Elem()

	// Assuming httpClient and baseEndpoint are the names of the fields in your services
	v.FieldByName("httpClient").Set(reflect.ValueOf(httpClient))
	v.FieldByName("baseEndpoint").Set(reflect.ValueOf(baseEndpoint))

	return t
}
