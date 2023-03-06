package checks

import (
	"strconv"

	"github.com/go-cmd/cmd"
)

func InGitDir() bool {
	app := "git"
	args := []string{"rev-parse", "--is-inside-git-dir"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout[0]

	if boolVal, err := strconv.ParseBool(out); err == nil {
		return boolVal == true
	}
	return false
}

func IsInWorktree() bool {
	app := "git"
	args := []string{"rev-parse", "--is-inside-work-tree"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout[0]

	if boolVal, err := strconv.ParseBool(out); err == nil {
		return boolVal == true
	}
	return false
}
