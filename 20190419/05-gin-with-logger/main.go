package main

import (
	"net/http"
	"os"
	"time"

	"gin-http/config"
	"gin-http/router"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Setting.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if config.Setting.Logs.Pretty {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: !config.Setting.Logs.Color,
			},
		)
	}

}

func main() {
	config.Load()
	setLogger()
	hander := router.Handler()

	s := &http.Server{
		Addr:         ":" + config.Setting.Server.Port,
		Handler:      hander,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info().Msg("start the api server")
	s.ListenAndServe()
}
