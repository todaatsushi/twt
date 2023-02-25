package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseBool(s string) bool {
	sClean := strings.Trim(s, "\n")
	sClean = strings.Trim(sClean, " ")
	boolVal, err := strconv.ParseBool(sClean)
	if err != nil {
		log.Fatal("Err when parsing bool from string", err, s)
		os.Exit(1)
	}
	return boolVal
}
