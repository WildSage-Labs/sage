package main

import (
	"github.com/WildSage-Labs/sage"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GitCommit Build variables
var (
	GitCommit     = "N/A"
	BinaryVersion = "development"
	BuildDate     = "N/A"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Debug().
		Str("Commit", GitCommit).
		Str("Version", BinaryVersion).
		Str("Build date", BuildDate).
		Msg("Sage is starting up")
	err, s := sage.NewSage()
	if err != nil {
		panic(err.Error())
	}
	s.Start()
}
