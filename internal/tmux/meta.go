package tmux

import "github.com/go-cmd/cmd"

func SendKeys(session string, toSend ...string) {
	app := "tmux"
	args := append([]string{"send-keys", "-t", session}, toSend...)
	c := cmd.NewCmd(app, args...)
	<-c.Start()
}
