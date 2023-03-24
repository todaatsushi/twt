package command

import "github.com/go-cmd/cmd"

func Run(app string, args ...string) (out, err []string) {
	c := cmd.NewCmd(app, args...)
	<-c.Start()

	output := c.Status().Stdout
	errors := c.Status().Stderr
	return output, errors
}
