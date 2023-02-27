package healthcheck

import (
	"log"
	"os"
)

func AssertTmux() {
	if inTmux := InTmuxSession(); !inTmux {
		log.Fatal("Not in tmux")
		os.Exit(1)
	}
}

func AssertWorkTree() {
	if inWorktree := IsBareRepo(); !inWorktree {
		log.Fatal("Not in worktree")
		os.Exit(1)
	}
}

func AssertReady() {
	AssertWorkTree()
	AssertTmux()
}
