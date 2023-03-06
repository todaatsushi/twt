package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/checks"
	"github.com/todaatsushi/twt/internal/git"
	"github.com/todaatsushi/twt/internal/tmux"
	"github.com/todaatsushi/twt/internal/utils"
)

var removeWorktree = &cobra.Command{
	Use:   "rm",
	Short: "Remove a git worktree, and optionally the linked branch.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		checks.AssertReady()

		branch := args[0]
		sessionName := utils.GenerateSessionNameFromBranch(branch)

		flags := cmd.Flags()
		deleteBranch, err := flags.GetBool("delete-branch")
		if err != nil {
			log.Fatal("Couldn't check delete-branch flag")
		}
		force, err := flags.GetBool("force")
		if err != nil {
			log.Fatal("Couldn't check force flag")
		}

		// Git cleanup
		branchExistsAndCheckedOut := git.HasBranch(branch, true)
		worktreeExists := git.HasWorktree(branch)
		if !branchExistsAndCheckedOut {
			log.Fatalf("Branch %s doesn't exist, or isn't checked out", branch)
		}
		if !worktreeExists {
			log.Fatalf("Can't delete worktree %s as it doesn't exist", branch)
		}
		git.RemoveWorktree(sessionName, branch, force, deleteBranch)

		// Tmux cleanup
		existingSessions := tmux.ListSessions(true)
		needToSwitchSession := tmux.HasSession(sessionName) && tmux.GetCurrentSessionName() == sessionName && len(existingSessions) > 1
		if needToSwitchSession {
			newSession := strings.ReplaceAll(existingSessions[1], "\"", "")
			if !tmux.HasSession(newSession) {
				log.Fatal("Session doesn't exist")
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
}
