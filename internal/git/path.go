package git

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/todaatsushi/twt/internal/checks"
	"github.com/todaatsushi/twt/internal/command"
)

func getGitDir() (string, error) {
	return os.Getwd()
}

func getBaseFromWorktree(runner command.Runner) (string, error) {
	app := "git"
	args := []string{"rev-parse", "--show-toplevel"}

	out, _ := runner.Run(app, args...)
	if len(out) == 0 {
		return "", errors.New("Couldn't get root git worktree dir - is this a git dir?")
	}
	return filepath.Dir(out[0]), nil
}

func GetBaseDir(runner command.Runner) (string, error) {
	if checks.IsInWorktree(runner) {
		return getBaseFromWorktree(runner)
	} else if checks.InGitDir(runner) {
		return getGitDir()
	}
	return "", errors.New("Not in worktree or git dir.")
}
