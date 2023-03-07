package tmux

import (
	"log"

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

func KillSession(name string) {
	app := "tmux"
	args := []string{"kill-session", "-t", name}
	c := cmd.NewCmd(app, args...)
	<-c.Start()
}

func GetCurrentSessionName() string {
	app := "tmux"
	args := []string{"display-message", "-p", "#S"}
	c := cmd.NewCmd(app, args...)
	<-c.Start()

	out := c.Status().Stdout
	if len(out) == 0 {
		log.Fatal("Couldn't fetch current tmux session name")
	}
	return out[0]

}

func ListSessions(justNames bool) []string {
	app := "tmux"
	args := []string{"list-sessions"}
	if justNames {
		fetchNameOpts := []string{"-F", "\"#{session_name}\""}
		args = append(args, fetchNameOpts...)
	}
	c := cmd.NewCmd(app, args...)
	<-c.Start()

	out := c.Status().Stdout
	if len(out) == 0 {
		log.Fatal("Couldn't fetch current tmux session name")
	}
	return out
}

func HasSession(name string) bool {
	app := "tmux"
	args := []string{"has-session", "-t", name}

	c := cmd.NewCmd(app, args...)
	<-c.Start()
	stderr := c.Status().Stderr

	return len(stderr) == 0

}
