package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "twt",
	Short: "Manage tmux sessions & windows based on Git worktrees.",
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		color.Red("Error when running cmd.")
	}
}
