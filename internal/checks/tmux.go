package checks

import (
	"log"
	"os"
)

func InTmuxSession() bool {
	isTmux, ok := os.LookupEnv("TMUX")
	if !ok {
		log.Fatal("Error when looking up TMUX env var")
	}
	return len(isTmux) > 0
}
