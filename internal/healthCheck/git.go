package healthcheck

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func isRepo() bool {
	isWorktree, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		log.Fatal("Error when checking worktree status", err)
		os.Exit(1)
	}
	return len(isWorktree) > 5
}

func IsBareRepo() bool {
	isRepo := isRepo()
	if !isRepo {
		return false
	}

	isWorktree, err := exec.Command("git", "rev-parse", "--is-bare-repository").Output()
	if err != nil {
		log.Fatal("Error when checking worktree status", err)
		os.Exit(1)
	}

	boolVal, err := strconv.ParseBool(string(isWorktree))
	if err != nil {
		log.Fatal("Error when parsing worktree val")
		os.Exit(1)
	}
	return boolVal
}
