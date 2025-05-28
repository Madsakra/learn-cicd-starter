package auth

import (
	"errors"
	"net/http"
	"testing"
)

func ErrMalformedHeader() error {
	return errors.New("malformed authorization header")
}

type TestFormat struct {
	name          string
	headers       http.Header
	expectedKey   string
	expectedError error
}

func TestGetApiKey(t *testing.T) {
	tests := []TestFormat{
		{
			name:          "No authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization header (no ApiKey)",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			expectedKey:   "",
			expectedError: ErrMalformedHeader(),
		},
		{
			name: "Malformed Authorization header (missing key)",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: ErrMalformedHeader(),
		},
		{
			name: "Proper ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			expectedKey:   "my-secret-key",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %q, got %q", tt.expectedError.Error(), err.Error())
			}
		})
	}

}
func unused() {
	// this function does nothing
	// and is called nowhere
}
