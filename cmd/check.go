package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/checks"
	"github.com/todaatsushi/twt/internal/utils"
)

func checkTmux() {
	color.White("\u270F Checking tmux status: must be in a session.")
	inTmux := checks.InTmuxSession()
	if inTmux {
		color.Green(" - in tmux session \u2713")
	} else {
		color.Red(" - not in tmux session \u2717")
	}
}

func checkGit() {
	color.White("\u270F Checking git status: must be in a worktree or .git dir.")
	inGitDir := checks.InGitDir()
	inWorktree := checks.IsInWorktree()
	gitValid := inWorktree || inGitDir

	if gitValid {
		color.Green(" - in either a git bare repo or worktree \u2713")
	} else {
		color.Red(" - not either a git bare repo or worktree \u2717")
	}
}

func checkCommonFilesDir() {
	// Common files dir is optional so in yellow when files not found
	color.White("\u270F Checking common files setup.")
	pathFuncs := []func() (string, error){utils.GetCommonFilesDirPath, utils.GetScriptsDirPath}
	checking := []string{"common", "scripts"}

	color.Cyan("\nChecking common dir exists in base dir and that it has a scripts dir.")
	for i, f := range pathFuncs {
		dir, _ := f()
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			color.Green(fmt.Sprintf(" - %s exists.", dir))
		} else {
			color.Yellow(fmt.Sprintf(" - %s doesn't exist.", checking[i]))
		}
	}

	color.Cyan("\nChecking that scripts dir contains a dir for valid commands, and that they have supported scripts.")
	cmdsToCheck := []string{"go", "rm"}
	scriptsToCheck := []string{"post.sh"}

	for i, c := range cmdsToCheck {
		dir, _ := utils.GetScriptsDirPathForCommand(c)
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			color.Green(fmt.Sprintf(" - %s exists.", dir))
		} else {
			color.Yellow(fmt.Sprintf(" - Dir for command %s doesn't exist.", cmdsToCheck[i]))
		}

		for _, fileName := range scriptsToCheck {
			filePath := fmt.Sprintf("%s/%s", dir, fileName)
			if _, err := os.Stat(filePath); !os.IsNotExist(err) {
				color.Green(fmt.Sprintf(" - Script %s exists for command %s.", fileName, c))
			} else {
				color.Yellow(fmt.Sprintf(" - Script %s missing for command %s.", fileName, c))
			}
		}
		fmt.Println()
	}
}

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
		color.White("\n\nKey conditions checks.")
		checkTmux()
		checkGit()

		color.White("\n\nOptional conditions checks.")
		checkCommonFilesDir()
	},
}

func init() {
	rootCmd.AddCommand(healthCheck)
}
