package auth

import (
	"errors"
	"net/http"
	"testing"
)

func Test_GetAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		headers  http.Header
		expected string
		err      error
	}{
		{
			name:     "no auth header",
			headers:  http.Header{},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			name:     "malformed auth header",
			headers:  http.Header{"Authorization": []string{"Bearer"}},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
		{
			name:     "valid auth header",
			headers:  http.Header{"Authorization": []string{"ApiKey abc123"}},
			expected: "abc123",
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetAPIKey(tt.headers)
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
			if err != nil && tt.err == nil {
				t.Errorf("unexpected error: %s", err)
			}
			if err == nil && tt.err != nil {
				t.Errorf("expected error: %s, got nil", tt.err)
			}
			if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error %s, got %s", tt.err, err)
			}
		})
	}
}
