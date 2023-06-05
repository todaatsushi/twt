package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/git"
	"github.com/todaatsushi/twt/internal/tmux"
	"github.com/todaatsushi/twt/internal/utils"
)

const NEW_DIR_PERM = 0700

var commonBase = &cobra.Command{
	Use:   "common",
	Short: "Configure twt utils.",
	Long:  "Create a new session or switch to the session starting in the common files dir.",
	Run: func(cmd *cobra.Command, args []string) {
		sessionName := "common"
		hasCommonSession := tmux.HasSession(sessionName)

		flags := cmd.Flags()
		removeSession, err := flags.GetBool("remove-session")
		if err != nil {
			color.Red("Error fetching the remove sesion flag")
			return
		}
		currentSession, err := tmux.GetCurrentSessionName()
		if err != nil && removeSession {
			color.Red("Can't remove current session")
		}

		if hasCommonSession {
			tmux.SwitchToSession(sessionName)
			if removeSession {
				tmux.KillSession(currentSession)
			}
			return
		}

		tmux.NewSession(sessionName)

		commonFilesDir, err := utils.GetCommonFilesDirPath()
		if err != nil {
			color.Red(fmt.Sprint(err))
			return
		}
		cdToCommonCommand := fmt.Sprintf("cd %s", commonFilesDir)
		tmux.SendKeys(sessionName, cdToCommonCommand, "Enter")
		tmux.SendKeys(sessionName, "clear", "Enter")

		tmux.SwitchToSession(sessionName)
		if removeSession {
			tmux.KillSession(currentSession)
		}
	},
}

var commonInit = &cobra.Command{
	Use:   "init",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Short: "Create common files in the bare repo, with templates for scripts.",
	Long: `
	'common' is a dir you can use in the bare repo dir that can house shared assets for
	your project, e.g. a common .env file and / or startup scripts.

	This command just creates the directory in the bare repo dir, and the 'scripts' folder
	which let's you define post worktree start scripts e.g. setting env vars.

	Currently supporting post command scripts only, named 'post.sh' within each command's
	dir in 'scripts'. Run 'twt check' to see where files should live.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		baseDir, err := git.GetBaseDir()
		if err != nil {
			color.Red(fmt.Sprint(err))
			return
		}

		color.Cyan("Setting up common file dir.\n\n")
		path := baseDir
		toCreate := []string{"common", "scripts"}

		// Create nested dirs
		for _, dir := range toCreate {
			path = fmt.Sprintf("%s/%s", path, dir)
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				color.Yellow(fmt.Sprintf("Dir %s exists - skipping.", dir))
				continue
			}
			color.Cyan(fmt.Sprintf("Creating %s.", path))

			if err := os.Mkdir(path, NEW_DIR_PERM); err != nil {
				color.Red(fmt.Sprintf("Dir %s couldn't be created: %s.", dir, err))
				return
			} else {
				color.Green(fmt.Sprintf("Successfully created %s in %s.", dir, path))
			}
		}

		color.Cyan("\n\nSetting up scripts dirs for each command.\n\n")
		// Command scripts for each command
		cmdDirs := []string{"go"}
		scripts := []string{"post.sh"}
		for _, dir := range cmdDirs {
			newPath := fmt.Sprintf("%s/%s", path, dir)

			if _, err := os.Stat(newPath); !os.IsNotExist(err) {
				color.Yellow(fmt.Sprintf("Dir %s exists - skipping.", dir))
				continue
			}
			color.Cyan(fmt.Sprintf("Creating %s.", newPath))

			if err := os.Mkdir(newPath, NEW_DIR_PERM); err != nil {
				color.Red(fmt.Sprintf("Dir %s couldn't be created: %s.", dir, err))
				return
			} else {
				color.Green(fmt.Sprintf("Successfully created %s in %s.", dir, newPath))
			}

			for _, fileName := range scripts {
				filePath := fmt.Sprintf("%s/%s", newPath, fileName)

				if _, err := os.Stat(filePath); !os.IsNotExist(err) {
					color.Yellow(fmt.Sprintf("File %s exists - skipping.", dir))
					continue
				}

				if f, err := os.Create(filePath); err != nil {
					color.Red(fmt.Sprintf("File %s couldn't be created: %s.", filePath, err))
					f.Close()
					return
				} else {
					color.Green(fmt.Sprintf("Successfully created %s", filePath))
					os.Chmod(filePath, 0700)
					f.WriteString("#!/bin/bash\n\n")
					f.WriteString("echo \"Enter your scripts here\"")
					f.Close()
				}
			}
		}
	},
}

func init() {
	// Register root
	rootCmd.AddCommand(commonBase)

	// Config topics
	commonBase.AddCommand(commonInit)

	commonBase.Flags().BoolP("remove-session", "r", false, "Remove current session (not worktree) after.")
}
