package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			expectedKey:   "valid_api_key",
			expectedError: nil,
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"InvalidHeader"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if apiKey != tt.expectedKey {
				t.Errorf("expected API key '%s', got '%s'", tt.expectedKey, apiKey)
			}
			if (err == nil && tt.expectedError != nil) || (err != nil && tt.expectedError == nil) || (err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("expected error '%v', got '%v'", tt.expectedError, err)
			}
		})
	}
}
