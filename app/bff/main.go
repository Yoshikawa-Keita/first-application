package main

import (
	"github.com/Yoshikawa-Keita/first-application/app/bff/router"
	"github.com/Yoshikawa-Keita/first-application/app/bff/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load configuration")
	}
	if config.Environment == "development" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
	e := router.NewRouter()
	e.Logger.Fatal(e.Start(config.HTTPBFFAddress))
}
