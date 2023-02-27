package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/todaatsushi/twt/internal/git"
	checks "github.com/todaatsushi/twt/internal/healthCheck"
	"github.com/todaatsushi/twt/internal/tmux"
	"github.com/todaatsushi/twt/internal/utils"
)

var goToWorktree = &cobra.Command{
	Use:   "go",
	Short: "Gets or creates a Tmux session from a given branch.",
	Long: `Given a branch name, either gets or creates a new Tmux session and creates
	/ switches to that branch within that session. If the session already exists,
	switches to it. Works for session names only if the session already exists.`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		checks.AssertReady()

		flags := cmd.Flags()
		switchSession, err := flags.GetBool("switch")
		if err != nil {
			log.Fatal("Couldn't fetch switch flag")
			os.Exit(1)
		}

		branch := args[0]
		sessionName := utils.SanitizeName(branch)
		isNewSession := !tmux.HasSession(sessionName)

		if !isNewSession {
			log.Printf("Session %s already exists.", sessionName)
			if switchSession {
				tmux.SwitchToSession(sessionName)
			}
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

		if switchSession {
			tmux.SwitchToSession(sessionName)
		}
	},
}

func init() {
	rootCmd.AddCommand(goToWorktree)

	goToWorktree.Flags().BoolP("switch", "s", false, "Switch to session after creation / retrieval.")
}
