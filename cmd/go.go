package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/todaatsushi/twt/internal/checks"
	"github.com/todaatsushi/twt/internal/command"
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
		runner := command.Terminal{}

		shouldCancel := checks.AssertReady(runner)
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
		branch, err = command.Validate(branch)
		if err != nil {
			color.Red(err.Error())
			return
		}

		removeSession, err := flags.GetBool("remove-session")
		if err != nil {
			color.Red("Error fetching the remove session flag")
			return
		}
		currentSession, err := tmux.GetCurrentSessionName()
		if err != nil && removeSession {
			color.Red("Can't remove current session.")
		}

		// Switch to session if exists
		sessionName := utils.GenerateSessionNameFromBranch(branch)
		isNewSession := !tmux.HasSession(sessionName)

		baseDir, err := git.GetBaseDir(runner)
		if err != nil {
			color.Red(fmt.Sprint(err))
			return
		}

		if !isNewSession {
			tmux.SwitchToSession(sessionName)
			if removeSession {
				tmux.KillSession(currentSession)
			}
			return
		}

		tmux.NewSession(sessionName)
		worktreeExists := git.HasWorktree(branch)
		if worktreeExists {
			changeDirCmd := fmt.Sprintf("cd %s/%s", baseDir, sessionName)
			tmux.SendKeys(sessionName, changeDirCmd, "Enter")
			tmux.SendKeys(sessionName, "clear", "Enter")
			tmux.SwitchToSession(sessionName)
			if removeSession {
				tmux.KillSession(currentSession)
			}
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

		// Execute post init scripts
		noScripts, err := flags.GetBool("no-scripts")
		if err != nil {
			color.Red("Couldn't fetch the run scripts flag")
			return
		}
		if !noScripts {
			utils.ExecuteScriptInSession(sessionName, "go", "post.sh")
		}

		if switchSession {
			tmux.SwitchToSession(sessionName)
			if removeSession {
				tmux.KillSession(currentSession)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(goToWorktree)

	goToWorktree.Flags().BoolP("switch", "s", false, "Switch to session after creation / retrieval.")
	goToWorktree.Flags().StringP("branch", "b", "", "Branch of which to create a new worktree + session.")
	goToWorktree.Flags().BoolP("no-scripts", "N", false, "Don't run any scripts in the common files dir if they exist for this command.")
	goToWorktree.Flags().BoolP("remove-session", "r", false, "Remove current session (not worktree) after.")
}
