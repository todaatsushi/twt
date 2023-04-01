package tmux

import (
	"github.com/todaatsushi/twt/internal/command"
)

func SendKeys(session string, toSend ...string) {
	args := append([]string{"send-keys", "-t", session}, toSend...)
	command.Run("tmux", args...)
}
