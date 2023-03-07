package checks

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func AssertTmux() error {
	if inTmux := InTmuxSession(); !inTmux {
		return errors.New("Not in tmux session")
	}
	return nil
}

func AssertGit() error {
	isWorktree := IsInWorktree()
	inGitDir := InGitDir()

	if valid := isWorktree || inGitDir; !valid {
		return errors.New("\u2717 Git status invalid - must be in a .git folder (worktree base) or inside a worktree")
	}
	return nil
}

func AssertReady() bool {
	shouldCancel := false

	gitErr := AssertGit()
	tmuxErr := AssertTmux()
	errs := [2]error{gitErr, tmuxErr}

	for _, err := range errs {
		if err != nil {
			msg := fmt.Sprint(err)
			color.Red(msg)
			shouldCancel = true
		}
	}
	return shouldCancel
}
