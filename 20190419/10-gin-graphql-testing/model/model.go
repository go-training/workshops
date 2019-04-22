package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/zerolog/log"
)

// DB connection
var DB *gorm.DB
var err error

// SetupDatabase connection
func SetupDatabase() error {
	DB, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal().Err(err).Msg("can't create db")
	}

	initialDB()

	return err
}

func initialDB() {
	DB.LogMode(true)
	if err := DB.AutoMigrate(&User{}).Error; err != nil {
		log.Fatal().Err(err).Msg("can't AutoMigrate db")
	}
}
