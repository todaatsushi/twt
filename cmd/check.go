package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	health "github.com/todaatsushi/twt/internal/healthCheck"
)

var healthCheck = &cobra.Command{
	Use:   "check",
	Short: "Check if twt is ready to be run in this shell",
	Long:  "twt needs to be run in a bare repo to use Git worktrees, and in a Tmux session.",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO - change println to be proper messaging
		isRepo := health.InTmuxSession()
		if isRepo {
			fmt.Println("In tmux session")
		} else {
			fmt.Println("Not in tmux session")
		}

		isBareRepo := health.IsBareRepo()
		if isBareRepo {
			fmt.Println("In bare repo")
		} else {
			fmt.Println("Not in bare repo")
		}

	},
}

func init() {
	rootCmd.AddCommand(healthCheck)
}
