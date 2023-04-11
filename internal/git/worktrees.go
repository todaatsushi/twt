package git

import (
	"github.com/todaatsushi/twt/internal/command"
)

func RemoveWorktree(name, branch string, force, deleteBranch bool) []string {
	app := "git"
	args := []string{"worktree", "remove", name}
	if force {
		args = append(args, "--force")
	}
	if _, errs := command.Run(app, args...); errs != nil {
		return errs
	}

	if deleteBranch {
		DeleteBranch(branch, force)
	}
	return nil
}
