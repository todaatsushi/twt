package git

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-cmd/cmd"
)

func GetBaseDir() string {
	app := "git"
	args := []string{"rev-parse", "--show-toplevel"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()

	out := c.Status().Stdout
	if len(out) == 0 {
		log.Fatal("Couldn't get root git worktree dir - is this a git dir?")
		os.Exit(1)
	}
	return filepath.Dir(out[0])
}
