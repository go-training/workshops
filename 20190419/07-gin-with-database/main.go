package main

import (
	"net/http"
	"os"
	"time"

	"gin-http-server/config"
	"gin-http-server/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
  setLogger()
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&router.User{})

	server := &http.Server{
		Addr:    ":" + config.Setting.Server.Port,
		Handler: router.Handler(db),
		// for upload data
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info().Msg("Server is running")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("can't start server")
	}
}

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
