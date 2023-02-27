package git

import (
	"fmt"

	"github.com/go-cmd/cmd"
)

func HasBranch(branch string) bool {
	app := "git"
	args := []string{"show-ref", fmt.Sprintf("refs/heads/%s", branch)}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout
	return len(out) > 0
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
