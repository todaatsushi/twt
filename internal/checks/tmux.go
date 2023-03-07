package checks

import (
	"os"
)

func InTmuxSession() bool {
	isTmux, ok := os.LookupEnv("TMUX")
	if !ok {
		return false
	}
	return len(isTmux) > 0
}
