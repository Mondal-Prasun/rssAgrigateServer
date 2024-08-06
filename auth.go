package main

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {insert apikey here}
func GetApiKeyFromHeader(headers http.Header) (apiKey string, err error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization key found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("nalformed header found")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header found")
	}
	return vals[1], nil
}
