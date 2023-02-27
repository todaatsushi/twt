package tmux

import (
	"github.com/go-cmd/cmd"
)

func SwitchToSession(name string) {
	app := "tmux"
	args := []string{"switch", "-t", name}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
}

func NewSession(branchName, cleanBranchName string) {
	app := "tmux"
	args := []string{"new-session", "-s", cleanBranchName, "-d"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
}

func HasSession(name string) bool {
	app := "tmux"
	args := []string{"has-session", "-t", name}

	c := cmd.NewCmd(app, args...)
	<-c.Start()
	stderr := c.Status().Stderr

	return len(stderr) == 0

}

func SendKeys(session string, toSend ...string) {
	app := "tmux"
	args := append([]string{"send-keys", "-t", session}, toSend...)
	c := cmd.NewCmd(app, args...)
	<-c.Start()
}
