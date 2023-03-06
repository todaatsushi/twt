package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/checks"
)

var healthCheck = &cobra.Command{
	Use:   "check",
	Short: "Check if twt is ready to be run in this shell",
	Long:  "twt needs to be run in a bare repo to use Git worktrees, and in a Tmux session.",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO - change println to be proper messaging
		isRepo := checks.InTmuxSession()
		if isRepo {
			fmt.Println("In tmux session")
		} else {
			fmt.Println("Not in tmux session")
		}

		inGitDir := checks.InGitDir()
		inWorktree := checks.IsInWorktree()

		gitValid := inWorktree || inGitDir
		if gitValid {
			fmt.Println("In worktree or git folder")
		} else {
			fmt.Printf("Not in worktree or git folder: worktree (%t) git folder (%t)", inWorktree, inGitDir)
		}

	},
}

func init() {
	rootCmd.AddCommand(healthCheck)
}
