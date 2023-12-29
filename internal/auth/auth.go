package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey - extracts api key from request header
// Example:
// Authorization: Apikey (insert key here)
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New(("no authentication info found"))
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New(("malformed authentication header"))
	}

	return vals[1], nil
}