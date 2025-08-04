package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/todaatsushi/twt/internal/command"
	"github.com/todaatsushi/twt/internal/git"
)

func GetCommonFilesDirPath() (string, error) {
	runner := command.Terminal{}
	baseDir, err := git.GetBaseDir(runner)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/common", baseDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("Common files dir doesn't exist.")
	}
	return path, nil
}

func GetScriptsDirPath() (string, error) {
	commonFilesDir, err := GetCommonFilesDirPath()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/scripts", commonFilesDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("Scripts dir doesn't exist.")
	}
	return path, nil
}

func GetScriptsDirPathForCommand(command string) (string, error) {
	scriptsDir, err := GetScriptsDirPath()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/%s", scriptsDir, command)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("Scripts dir doesn't exist.")
	}
	return path, nil
}
