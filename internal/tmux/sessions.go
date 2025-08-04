package tmux

import (
	"errors"

	"github.com/todaatsushi/twt/internal/command"
)

func SwitchToSession(runner command.Runner, name string) {
	runner.Run("tmux", "switch", "-t", name)
}

func NewSession(runner command.Runner, cleanBranchName string) {
	runner.Run("tmux", "new-session", "-s", cleanBranchName, "-d")
}

func KillSession(runner command.Runner, name string) {
	runner.Run("tmux", "kill-session", "-t", name)
}

func GetCurrentSessionName(runner command.Runner) (string, error) {
	out, _ := runner.Run("tmux", "display-message", "-p", "#S")
	if len(out) == 0 {
		return "", errors.New("Couldn't fetch current tmux session name")
	}
	return out[0], nil

}

func ListSessions(runner command.Runner, justNames bool) ([]string, error) {
	args := []string{"list-sessions"}
	if justNames {
		fetchNameOpts := []string{"-F", "\"#{session_name}\""}
		args = append(args, fetchNameOpts...)
	}
	out, _ := runner.Run("tmux", args...)

	if len(out) == 0 {
		return []string{}, errors.New("Couldn't fetch current tmux session name")
	}
	return out, nil
}

func HasSession(runner command.Runner, name string) bool {
	_, stderr := runner.Run("tmux", "has-session", "-t", name)
	return len(stderr) == 0

}
