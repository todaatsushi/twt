package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-cmd/cmd"
)

func runAndStreamCmd(sessionName, command string) {
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	var envCmd *cmd.Cmd
	if sessionName == "" {
		envCmd = cmd.NewCmdOptions(cmdOptions, command)
	} else {
		envCmd = cmd.NewCmdOptions(cmdOptions, "tmux", "send-keys", "-t", sessionName, command, "Enter")
	}

	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
				fmt.Println(line)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	<-envCmd.Start()
	<-doneChan
}

func ExecuteScriptInSession(sessionName, command, scriptName string) error {
	commandScriptsDir, err := GetScriptsDirPathForCommand(command)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", commandScriptsDir, scriptName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		msg := fmt.Sprintf("%s doesn't exist for %s. Create %s in %s to run.", scriptName, command, scriptName, path)
		return errors.New(msg)
	}

	runAndStreamCmd(sessionName, path)
	return nil
}
