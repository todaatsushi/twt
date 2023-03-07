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
	out := c.Status().Stdout

	if len(out) > 0 {
		insideGitDir := out[0]
		if boolVal, err := strconv.ParseBool(insideGitDir); err == nil {
			return boolVal == true
		}
	}
	return false
}

func IsInWorktree() bool {
	app := "git"
	args := []string{"rev-parse", "--is-inside-work-tree"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout

	if len(out) > 0 {
		insideGitWorktree := out[0]
		if boolVal, err := strconv.ParseBool(insideGitWorktree); err == nil {
			return boolVal == true
		}
	}
	return false
}
