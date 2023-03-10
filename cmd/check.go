package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/checks"
)

var healthCheck = &cobra.Command{
	Use:   "check",
	Short: "Check if twt is ready to be run in this shell",
	Long:  "twt needs to be run in a bare repo to use Git worktrees, and in a Tmux session.",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		color.White("\u270F Checking tmux status: must be in a session.")
		inTmux := checks.InTmuxSession()
		if inTmux {
			color.Green(" - in tmux session \u2713")
		} else {
			color.Red(" - not in tmux session \u2717")
		}

		color.White("\u270F Checking git status: must be in a worktree or .git dir.")
		inGitDir := checks.InGitDir()
		inWorktree := checks.IsInWorktree()
		gitValid := inWorktree || inGitDir

		if gitValid {
			color.Green(" - in either a git bare repo or worktree \u2713")
		} else {
			color.Red(" - not either a git bare repo or worktree \u2717")
		}

	},
}

func init() {
	rootCmd.AddCommand(healthCheck)
}
