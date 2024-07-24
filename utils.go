package redmine

import (
	"fmt"
	"net/url"
	"strings"
)

func removeAfterComma(input string) string {
	// Find the index of the comma
	commaIndex := strings.Index(input, ",")

	// If there's no comma, return the original string
	if commaIndex == -1 {
		return input
	}

	// Slice the string up to the comma
	return input[:commaIndex]
}

func mapToQueryString(params map[string]string) string {
	var queryParts []string
	for key, value := range params {
		escapedKey := url.QueryEscape(key)
		escapedValue := url.QueryEscape(value)
		queryParts = append(queryParts, fmt.Sprintf("%s=%s", escapedKey, escapedValue))
	}
	return strings.Join(queryParts, "&")
}
