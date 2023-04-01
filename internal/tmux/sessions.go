package tmux

import (
	"errors"

	"github.com/todaatsushi/twt/internal/command"
)

func SwitchToSession(name string) {
	command.Run("tmux", "switch", "-t", name)
}

func NewSession(cleanBranchName string) {
	command.Run("tmux", "new-session", "-s", cleanBranchName, "-d")
}

func KillSession(name string) {
	command.Run("tmux", "kill-session", "-t", name)
}

func GetCurrentSessionName() (string, error) {
	out, _ := command.Run("tmux", "display-message", "-p", "#S")
	if len(out) == 0 {
		return "", errors.New("Couldn't fetch current tmux session name")
	}
	return out[0], nil

}

func ListSessions(justNames bool) ([]string, error) {
	args := []string{"list-sessions"}
	if justNames {
		fetchNameOpts := []string{"-F", "\"#{session_name}\""}
		args = append(args, fetchNameOpts...)
	}
	out, _ := command.Run("tmux", args...)

	if len(out) == 0 {
		return []string{}, errors.New("Couldn't fetch current tmux session name")
	}
	return out, nil
}

func HasSession(name string) bool {
	_, stderr := command.Run("tmux", "has-session", "-t", name)
	return len(stderr) == 0

}
