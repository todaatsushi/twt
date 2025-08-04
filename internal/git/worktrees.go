package git

import (
	"github.com/todaatsushi/twt/internal/command"
)

func RemoveWorktree(runner command.Runner, name, branch string, force, deleteBranch bool) []string {
	app := "git"
	args := []string{"worktree", "remove", name}
	if force {
		args = append(args, "--force")
	}
	_, errs := runner.Run(app, args...)
	if errs != nil {
		if len(errs) != 0 {
			return errs
		}
	}

	if deleteBranch {
		DeleteBranch(branch, force)
	}
	return nil
}
