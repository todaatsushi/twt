package checks

import (
	"strconv"

	"github.com/todaatsushi/twt/internal/command"
)

func InGitDir(runner command.Runner) bool {
	out, _ := runner.Run("git", "rev-parse", "--is-inside-git-dir")

	if len(out) > 0 {
		insideGitDir := out[0]
		if boolVal, err := strconv.ParseBool(insideGitDir); err == nil {
			return boolVal == true
		}
	}
	return false
}

func IsInWorktree(runner command.Runner) bool {
	out, _ := runner.Run("git", "rev-parse", "--is-inside-work-tree")

	if len(out) > 0 {
		insideGitWorktree := out[0]
		if boolVal, err := strconv.ParseBool(insideGitWorktree); err == nil {
			return boolVal == true
		}
	}
	return false
}

func IsUsingBareRepo(runner command.Runner) bool {
	out, _ := runner.Run("bash", "-c", "git worktree list | grep \\(bare\\)")
	return len(out) > 0
}
