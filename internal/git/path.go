package git

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/go-cmd/cmd"
	"github.com/todaatsushi/twt/internal/checks"
)

func GetBaseDir() (string, error) {
	app := "git"
	args := []string{"rev-parse", "--show-toplevel"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()

	if checks.IsInWorktree() {
		out := c.Status().Stdout
		if len(out) == 0 {
			log.Fatal("Couldn't get root git worktree dir - is this a git dir?")
		}
		return filepath.Dir(out[0]), nil
	}
	if checks.InGitDir() {
		dir, err := os.Getwd()
		return dir, err
	}
	return "", errors.New("Couldn't fetch git base dir. Base dir is returned when command is run inside a worktree or the git dir (ie. the base of the worktree).")
}
