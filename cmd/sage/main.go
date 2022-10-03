package main

import (
	"fmt"
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

//TODO: Read

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	fmt.Println("Hello, world! I am a wise sage that stares into cosmos")
	err, sage := sage.NewSage()
	if err != nil {
		panic(err.Error())
	}
	sage.Start()
}
