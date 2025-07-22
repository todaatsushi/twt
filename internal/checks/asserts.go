package checks

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/todaatsushi/twt/internal/command"
)

func AssertTmux() error {
	if inTmux := InTmuxSession(); !inTmux {
		return errors.New("\u2717 Not in tmux session")
	}
	return nil
}

func AssertGit(runner command.Runner) error {
	isWorktree := IsInWorktree(runner)
	inGitDir := InGitDir(runner)

	inGit := isWorktree || inGitDir
	usingBareRepo := IsUsingBareRepo(runner)

	if valid := inGit && usingBareRepo; !valid {
		return errors.New("\u2717 Git status invalid - must be in a .git folder (worktree base) or inside a worktree, and in a bare repository.")
	}
	return nil
}

func AssertReady() bool {
	// Init here instead of return in the loop to show all messages
	shouldCancel := false
	runner := command.Terminal{}

	gitErr := AssertGit(runner)
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
