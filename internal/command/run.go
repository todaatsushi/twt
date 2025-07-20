package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-cmd/cmd"
)

type Runner interface {
	Run(app string, args ...string) (out, err []string)
}

func Run(app string, args ...string) (out, err []string) {
	// To be deprecated in favor of running in the interface
	c := cmd.NewCmd(app, args...)
	<-c.Start()

	output := c.Status().Stdout
	errors := c.Status().Stderr
	return output, errors
}

func Validate(branchName string) (string, error) {
	hasSpaces := strings.Contains(branchName, " ")
	hasColon := strings.Contains(branchName, ";")
	hasNewline := strings.Contains(branchName, "\n")

	errMessage := "Error: illegal character(s):"
	if hasSpaces {
		errMessage = fmt.Sprintf("%s \"%s\"", errMessage, " ")
	}
	if hasColon {
		errMessage = fmt.Sprintf("%s \"%s\"", errMessage, ";")
	}
	if hasNewline {
		errMessage = fmt.Sprintf("%s \"%s\"", errMessage, "\\n")
	}

	if hasSpaces || hasColon || hasNewline {
		return "", errors.New(errMessage)
	}
	return branchName, nil
}
