package git

import (
	"fmt"
	"strings"

	"github.com/go-cmd/cmd"
)

func HasBranch(branch string, checkedOut bool) bool {
	app := "git"
	args := []string{"branch"}
	if !checkedOut {
		args = []string{"show-ref", fmt.Sprintf("refs/heads/%s", branch)}
	}

	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout

	if !checkedOut {
		return len(out) > 0
	}

	checkedOutBranches := make(map[string]bool)
	var clean string
	for i := 0; i < len(out); i++ {
		s := out[i]
		if strings.Contains(s, "+") || strings.Contains(s, "*") {
			clean = strings.Replace(s, "+", "", 1)
			clean = strings.Replace(clean, "*", "", 1)
			clean = strings.Replace(clean, " ", "", 1)
			checkedOutBranches[clean] = true
		}
	}
	_, ok := checkedOutBranches[branch]
	return ok
}

func HasWorktree(branch string) bool {
	grepCmd := fmt.Sprintf("git worktree list | grep %s", branch)
	c := cmd.NewCmd("bash", "-c", grepCmd)
	<-c.Start()
	out := c.Status().Stdout
	return len(out) > 0
}

func DeleteBranch(branch string, force bool) {
	app := "git"

	deleteFlag := "-d"
	if force {
		deleteFlag = "-D"
	}

	args := []string{"branch", deleteFlag, branch}

	c := cmd.NewCmd(app, args...)
	<-c.Start()
}
