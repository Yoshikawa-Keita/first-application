// Server/main.go

package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"yoshikawa-keita-first-application/app/server/router"
	"yoshikawa-keita-first-application/app/server/util"
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
	e.Logger.Fatal(e.Start(config.HTTPServerAddress))
}
