package main

import (
	"fmt"
	"os"

	"github.com/WildSage-Labs/sage/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	fmt.Println("Hello, world! I am a wise sage that stares into cosmos")
	err, sage := internal.NewSage()
	if err != nil {
		panic(err.Error())
	}
	sage.Start()
}
