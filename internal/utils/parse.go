package utils

import (
	"log"
	"os"
	"strconv"
)

func ParseBool(s string) bool {
	boolVal, err := strconv.ParseBool(SanatizeString(s))
	if err != nil {
		log.Fatal("Err when parsing bool from string", err, s)
		os.Exit(1)
	}
	return boolVal
}
