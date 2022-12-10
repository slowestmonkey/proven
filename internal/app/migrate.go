package app

import (
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func init() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	connection := viper.GetString(`database.connection`)

	var (
		attempts = _defaultAttempts
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New(
			"file://migrations",
			connection,
		)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		panic(err)
	}

	m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}
	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}
	log.Printf("Migrate: up success")
}
