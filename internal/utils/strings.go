package utils

import "strings"

func SanatizeString(s string) string {
	cleanStr := strings.Trim(s, "\n")
	return cleanStr
}
