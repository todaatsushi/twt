package healthcheck

import (
	"log"
	"os"
	"os/exec"

	"github.com/todaatsushi/twt/internal/utils"
)

func isRepo() bool {
	isRepo, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		log.Fatal("Error when checking repo status", err)
		os.Exit(1)
	}

	isRepoStr := utils.SanatizeString(string(isRepo))
	boolVal := utils.ParseBool(isRepoStr)
	return boolVal
}

func IsBareRepo() bool {
	// TODO - find way to verify in bare repo cleanly
	// isRepo := isRepo()
	// if !isRepo {
	// 	return false
	// }

	// isWorktree, err := exec.Command("git", "rev-parse", "--is-bare-repository").Output()
	// if err != nil {
	// 	log.Fatal("Error when checking worktree status", err)
	// 	os.Exit(1)
	// }

	// boolVal := utils.ParseBool(string(isWorktree))
	// return boolVal
	return isRepo()
}
