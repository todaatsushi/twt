package healthcheck

import (
	"log"
	"os"
)

func InTmuxSession() bool {
	isTmux, ok := os.LookupEnv("TMUX")
	if !ok {
		log.Fatal("Error when looking up TMUX env var")
		os.Exit(1)
	}
	return len(isTmux) > 0
}
