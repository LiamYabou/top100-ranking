package logger

// Switch the level according with the different environment.

import (
	"github.com/LiamYabou/top100-ranking/pkg/variable"
	log "github.com/sirupsen/logrus"
)

func switchError(entry *log.Entry, msg string) {
	switch variable.Env {
	case "development":
		entry.Panic(msg)
	case "staging":
		entry.Panic(msg)
	case "production":
		entry.Fatal(msg)
	}
}
