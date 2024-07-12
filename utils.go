package redmine

import "strings"

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
