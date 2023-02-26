package utils

import "strings"

func SanitizeName(branchName string) string {
	return strings.Replace(branchName, "/", "__", -1)
}
