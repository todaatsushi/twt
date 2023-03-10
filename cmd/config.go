package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/todaatsushi/twt/internal/git"
)

const NEW_DIR_PERM = 0700

var config = &cobra.Command{
	Use:   "config",
	Short: "Configure twt utils.",
	Long: `
	Customise twt to your individual needs:

	- common: shared resources across sessions + worktrees e.g. post init scripts.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Red("Please specify a sub command to configure.")
	},
}

var common = &cobra.Command{
	Use:   "common",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Short: "Shared resources across sessions + worktrees e.g. post init scripts.",
	Long: `
	'common' is a dir you can use in the bare repo dir that can house shared assets for
	your project, e.g. a common .env file and / or startup scripts.

	This command just creates the directory in the bare repo dir, and the 'scripts' folder
	which let's you define post worktree start scripts e.g. setting env vars.
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
		cmdDirs := []string{"go", "rm"}
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

				if _, err := os.Create(filePath); err != nil {
					color.Red(fmt.Sprintf("File %s couldn't be created: %s.", filePath, err))
					return
				} else {
					color.Green(fmt.Sprintf("Successfully created %s", filePath))
				}
			}
		}
	},
}

func init() {
	// Register root
	rootCmd.AddCommand(config)

	// Config topics
	config.AddCommand(common)
}
