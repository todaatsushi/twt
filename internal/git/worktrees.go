package git

import (
	"github.com/todaatsushi/twt/internal/command"
)

func RemoveWorktree(name, branch string, force, deleteBranch bool) {
	app := "git"
	args := []string{"worktree", "remove", name}
	if force {
		args = append(args, "--force")
	}
	command.Run(app, args...)

	if deleteBranch {
		DeleteBranch(branch, force)
	}
}
