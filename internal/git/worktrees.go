package git

import "github.com/go-cmd/cmd"

func RemoveWorktree(worktree string, force bool) {
	app := "git"
	args := []string{"worktree", "remove", worktree}

	if force {
		args = append(args, "--force")
	}

	c := cmd.NewCmd(app, args...)
	<-c.Start()
}
