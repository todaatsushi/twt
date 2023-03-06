package checks

import (
	"log"
)

func AssertTmux() {
	if inTmux := InTmuxSession(); !inTmux {
		log.Fatal("Not in tmux")
	}
}

func AssertGit() {
	isWorktree := IsInWorktree()
	inGitDir := InGitDir()

	if valid := isWorktree || inGitDir; !valid {
		log.Fatal("Git status invalid - must be in a .git folder (worktree base) or inside a worktree")
	}

}

func AssertReady() {
	AssertGit()
	AssertTmux()
}
