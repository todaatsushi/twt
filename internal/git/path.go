package git

import (
	"errors"
	"path/filepath"

	"github.com/go-cmd/cmd"
	"github.com/todaatsushi/twt/internal/checks"
)

func GetBaseDir() (string, error) {
	app := "git"
	args := []string{"rev-parse", "--show-toplevel"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout
	if len(out) == 0 {
		return "", errors.New("Couldn't get root git worktree dir - is this a git dir?")
	}

	if checks.IsInWorktree() {
		return filepath.Dir(out[0]), nil
	}
	if checks.InGitDir() {
		return out[0], nil
	}
	return "", errors.New("Couldn't fetch git base dir. Base dir is returned when command is run inside a worktree or the git dir (ie. the base of the worktree).")
}
