package git

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/go-cmd/cmd"
	"github.com/todaatsushi/twt/internal/checks"
)

func getGitDir() (string, error) {
	return os.Getwd()
}

func getBaseFromWorktree() (string, error) {
	app := "git"
	args := []string{"rev-parse", "--show-toplevel"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout

	if len(out) == 0 {
		return "", errors.New("Couldn't get root git worktree dir - is this a git dir?")
	}
	return filepath.Dir(out[0]), nil
}

func GetBaseDir() (string, error) {
	if checks.IsInWorktree() {
		return getBaseFromWorktree()
	} else if checks.InGitDir() {
		return getGitDir()
	}
	return "", errors.New("Not in worktree or git dir.")
}
