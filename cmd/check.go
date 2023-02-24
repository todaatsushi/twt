package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	health "github.com/todaatsushi/twt/internal/healthCheck"
)

var healthCheck = &cobra.Command{
	Use: "check",
	Run: func(cmd *cobra.Command, args []string) {
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
