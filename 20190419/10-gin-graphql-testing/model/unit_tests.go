package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rs/zerolog/log"
	"gopkg.in/testfixtures.v2"
)

var fixtures *testfixtures.Context

// MainTest a reusable TestMain(..) function for unit tests that need to use a
// test database. Creates the test database, and sets necessary settings.
func MainTest(m *testing.M, root string) {
	var err error
	if err := SetupDatabase(); err != nil {
		log.Fatal().Err(err).Msg("can't connect database")
	}
	fixturesDir := filepath.Join(root, "model", "fixtures")

	testfixtures.SkipDatabaseNameCheck(true)
	fixtures, err = testfixtures.NewFolder(DB.DB(), &testfixtures.SQLite{}, fixturesDir)
	if err != nil {
		log.Fatal().Err(err).Msg("can't connect database")
	}

	os.Exit(m.Run())
}

func PrepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatal().Err(err).Msg("can't load fixtures")
	}
}
