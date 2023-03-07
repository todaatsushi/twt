package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "twt",
	Short: "Manage tmux sessions & windows based on Git worktrees",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("Error when running root cmd.")
	}
}
