package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/checks"
	"github.com/todaatsushi/twt/internal/git"
)

var healthCheck = &cobra.Command{
	Use:   "check",
	Short: "Check if twt is ready to be run in this shell.",
	Long: `
	twt needs to be run in a bare repo to use Git worktrees, and in a Tmux session.

	Check also optional usage of common files, which can be used to configure a shared
	state between worktrees.
	`,
	Args: cobra.MatchAll(cobra.ExactArgs(0)),
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

		color.White("\u270F Checking common files setup.")
		baseDir, err := git.GetBaseDir()
		if err != nil {
			color.Red(fmt.Sprint(err))
			return
		}

		path := baseDir
		toCheck := []string{"common", "scripts"}
		for _, dir := range toCheck {
			path = fmt.Sprintf("%s/%s", path, dir)
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				color.Green(fmt.Sprintf(" - %s exists in %s.", dir, path))
			} else {
				color.Yellow(fmt.Sprintf(" - %s doesn't exist. Expected in %s.", dir, path))
			}
		}

		cmdsToCheck := []string{"go", "rm"}
		scriptsToCheck := []string{"post.sh"}

		for _, dir := range cmdsToCheck {
			newPath := fmt.Sprintf("%s/%s", path, dir)
			if _, err := os.Stat(newPath); !os.IsNotExist(err) {
				color.Green(fmt.Sprintf(" - %s exists in %s.", dir, newPath))
			} else {
				color.Yellow(fmt.Sprintf(" - Dir for command %s doesn't exist. Expected in %s.", dir, newPath))
			}

			for _, fileName := range scriptsToCheck {
				filePath := fmt.Sprintf("%s/%s", newPath, fileName)
				if _, err := os.Stat(filePath); !os.IsNotExist(err) {
					color.Green(fmt.Sprintf(" - %s exists in %s.", dir, filePath))
				} else {
					color.Yellow(fmt.Sprintf(" - Script %s missing for command %s.", fileName, newPath))
				}

			}
		}

	},
}

func init() {
	rootCmd.AddCommand(healthCheck)
}
