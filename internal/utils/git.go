package utils

import "strings"

func GenerateSessionNameFromBranch(branchName string) string {
	return strings.Replace(branchName, "/", "__", -1)
}
