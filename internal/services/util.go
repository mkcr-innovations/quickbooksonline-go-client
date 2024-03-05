package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
