package git

import "github.com/go-cmd/cmd"

func RemoveWorktree(name, branch string, force, deleteBranch bool) {
	app := "git"
	args := []string{"worktree", "remove", name}
	if force {
		args = append(args, "--force")
	}
	c := cmd.NewCmd(app, args...)
	<-c.Start()

	if deleteBranch {
		DeleteBranch(branch, force)
	}
}
