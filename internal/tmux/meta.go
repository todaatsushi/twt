package tmux

import (
	"github.com/todaatsushi/twt/internal/command"
)

func SendKeys(runner command.Runner, session string, toSend ...string) {
	args := append([]string{"send-keys", "-t", session}, toSend...)
	runner.Run("tmux", args...)
}
