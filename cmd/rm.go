package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/checks"
	"github.com/todaatsushi/twt/internal/git"
	"github.com/todaatsushi/twt/internal/tmux"
	"github.com/todaatsushi/twt/internal/utils"
)

var removeWorktree = &cobra.Command{
	Use:   "rm",
	Short: "Remove a git worktree, tmux session, and optionally the linked branch.",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		shouldCancel := checks.AssertReady()
		if shouldCancel {
			color.Red("Error when trying to run command, aborting.")
			return
		}

		flags := cmd.Flags()
		branch, err := flags.GetString("branch")
		if err != nil || branch == "" {
			color.Red("Couldn't fetch target branch.")
			return
		}

		sessionName := utils.GenerateSessionNameFromBranch(branch)

		deleteBranch, err := flags.GetBool("delete-branch")
		if err != nil {
			color.Red("Couldn't check delete-branch flag")
			return
		}
		force, err := flags.GetBool("force")
		if err != nil {
			color.Red("Couldn't check force flag")
			return
		}

		// Git cleanup
		branchExistsAndCheckedOut := git.HasBranch(branch, true)
		worktreeExists := git.HasWorktree(branch)
		if !branchExistsAndCheckedOut {
			color.Red("Branch %s doesn't exist, or isn't checked out", branch)
			return
		}
		if !worktreeExists {
			color.Red("Can't delete worktree %s as it doesn't exist", branch)
			return
		}
		git.RemoveWorktree(sessionName, branch, force, deleteBranch)

		// Tmux cleanup
		existingSessions, err := tmux.ListSessions(true)
		if err != nil {
			color.Red(fmt.Sprint(err))
			return
		}
		currentSession, err := tmux.GetCurrentSessionName()
		if err != nil {
			color.Red(fmt.Sprint(err))
			return
		}
		possibleDestinations := []string{}
		for _, session := range existingSessions {
			if session != currentSession {
				possibleDestinations = append(possibleDestinations, session)
			}
		}

		needToSwitchSession := tmux.HasSession(sessionName) && currentSession == sessionName && len(possibleDestinations) > 0
		if needToSwitchSession {
			newSession := strings.ReplaceAll(possibleDestinations[0], "\"", "")
			if !tmux.HasSession(newSession) {
				color.Red("Session doesn't exist")
				return
			}
			tmux.SwitchToSession(newSession)
		}
		tmux.KillSession(sessionName)
	},
}

func init() {
	rootCmd.AddCommand(removeWorktree)
	removeWorktree.Flags().BoolP("delete-branch", "d", false, "Remove branch as well as the worktree")
	removeWorktree.Flags().BoolP("force", "f", false, "Delete the worktree &| branch regardless of unstaged files")
	removeWorktree.Flags().StringP("branch", "b", "", "Branch of which to create a new worktree + session.")
}
