package healthcheck

import (
	"github.com/go-cmd/cmd"
	"github.com/todaatsushi/twt/internal/utils"
)

func isRepo() bool {
	app := "git"
	args := []string{"rev-parse", "--is-inside-work-tree"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
	out := c.Status().Stdout[0]
	boolVal := utils.ParseBool(out)
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
