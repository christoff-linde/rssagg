package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP request
// Example:
// Authorization: ApiKey <api_key>
func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("No authentication info found")
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		return "", errors.New("Malformed auth header")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth header")
	}
	return values[1], nil
}
