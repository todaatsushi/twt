package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/todaatsushi/twt/internal/checks"
	"github.com/todaatsushi/twt/internal/git"
	"github.com/todaatsushi/twt/internal/tmux"
	"github.com/todaatsushi/twt/internal/utils"
)

var goToWorktree = &cobra.Command{
	Use:   "go",
	Short: "Gets or creates a tmux session from a given branch.",
	Long: `Given a branch name, either gets or creates a new Tmux session and creates
	/ switches to that branch within that session.

	If the session already exists, switches to it regardless of if a git worktree exists
	or not. If this isn't desired, rename / delete the existing session.

	Also switches to a new session if a worktree exists (ie. the branch is checked out).
	`,
	Args: cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		shouldCancel := checks.AssertReady()
		if shouldCancel {
			color.Red("Error when trying to run command, aborting.")
			return
		}

		flags := cmd.Flags()
		switchSession, err := flags.GetBool("switch")
		if err != nil {
			color.Red("Couldn't fetch switch session flag.")
			return
		}
		branch, err := flags.GetString("branch")
		if err != nil || branch == "" {
			color.Red("Couldn't fetch target branch.")
			return
		}

		// Switch to session if exists
		sessionName := utils.GenerateSessionNameFromBranch(branch)
		isNewSession := !tmux.HasSession(sessionName)

		baseDir, err := git.GetBaseDir()
		if err != nil {
			color.Red(fmt.Sprint(err))
			return
		}

		if !isNewSession {
			tmux.SwitchToSession(sessionName)
			return
		}

		tmux.NewSession(sessionName)
		worktreeExists := git.HasWorktree(branch)
		if worktreeExists {
			changeDirCmd := fmt.Sprintf("cd %s/%s", baseDir, sessionName)
			tmux.SendKeys(sessionName, changeDirCmd, "Enter")
			tmux.SendKeys(sessionName, "clear", "Enter")
			tmux.SwitchToSession(sessionName)
			return
		}

		// Change to worktree base to create the worktree here
		backToBaseDirCmd := fmt.Sprintf("cd %s", baseDir)
		tmux.SendKeys(sessionName, backToBaseDirCmd, "Enter")

		branchIsNew := !git.HasBranch(branch, false)
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

		if switchSession {
			tmux.SwitchToSession(sessionName)
		}
	},
}

func init() {
	rootCmd.AddCommand(goToWorktree)

	goToWorktree.Flags().BoolP("switch", "s", false, "Switch to session after creation / retrieval.")
	goToWorktree.Flags().StringP("branch", "b", "", "Branch of which to create a new worktree + session.")
}
