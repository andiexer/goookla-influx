package utils

import "strings"

func RemoveQuotesFromString(s string) string {
	return strings.ReplaceAll(s, "\"", "")
}
