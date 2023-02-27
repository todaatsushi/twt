package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/todaatsushi/twt/internal/git"
	checks "github.com/todaatsushi/twt/internal/healthCheck"
	"github.com/todaatsushi/twt/internal/tmux"
	"github.com/todaatsushi/twt/internal/utils"
)

var goTree = &cobra.Command{
	Use:  "go",
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		checks.AssertReady()

		branch := args[0]
		sessionName := utils.SanitizeName(branch)
		isNewSession := !tmux.HasSession(sessionName)

		if !isNewSession {
			tmux.SwitchToSession(sessionName)
			return
		}

		tmux.NewSession(branch, sessionName)
		worktreeExists := git.HasWorktree(branch)
		if worktreeExists {
			baseDir := git.GetBaseDir()
			changeDirCmd := fmt.Sprintf("cd %s/%s", baseDir, sessionName)
			tmux.SendKeys(sessionName, changeDirCmd, "Enter")
			tmux.SendKeys(sessionName, "clear", "Enter")
			tmux.SwitchToSession(sessionName)
			return
		}

		// Change to worktree base to create the worktree here
		baseDir := git.GetBaseDir()
		backToBaseDirCmd := fmt.Sprintf("cd %s", baseDir)
		tmux.SendKeys(sessionName, backToBaseDirCmd, "Enter")

		branchIsNew := !git.HasBranch(branch)
		if branchIsNew {
			newWorktreeCmd := fmt.Sprintf("git worktree add %s -b %s", sessionName, branch)
			tmux.SendKeys(sessionName, newWorktreeCmd, "Enter")
		} else {
			newWorktreeCmd := fmt.Sprintf("git worktree add %s %s", sessionName, branch)
			tmux.SendKeys(sessionName, newWorktreeCmd, "Enter")
		}

		changeToNewTreeCmd := fmt.Sprintf("cd %s", sessionName)
		tmux.SendKeys(sessionName, changeToNewTreeCmd, "Enter")
		tmux.SendKeys(sessionName, "clear", "Enter")
	},
}

func init() {
	rootCmd.AddCommand(goTree)
}
