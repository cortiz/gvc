package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var Log zerolog.Logger

func init() {
	initLogger()
}

func initLogger() {
	gvcDebug := os.Getenv("GVC_DEBUG")
	Log = log.Output(zerolog.ConsoleWriter{TimeFormat: time.RFC3339, Out: os.Stdout, NoColor: false})
	if gvcDebug == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
