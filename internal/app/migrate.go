package app

// import (
// 	"errors"
// 	"log"
// 	"time"

// 	"github.com/golang-migrate/migrate/v4"
// )

// const (
// 	_defaultAttempts = 20
// 	_defaultTimeout  = time.Second
// )

// func Migrate(connection string) migrate.Migrate {
// 	// connection := viper.GetString(`database.connection`)

// 	var (
// 		attempts = _defaultAttempts
// 		m        *migrate.Migrate
// 		err      error
// 	)

// 	for attempts > 0 {
// 		m, err = migrate.New(
// 			"file://../../db/migrations",
// 			connection,
// 		)
// 		if err == nil {
// 			break
// 		}

// 		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
// 		time.Sleep(_defaultTimeout)
// 		attempts--
// 	}

// 	if err != nil {
// 		panic(err)
// 	}

// 	m.Up()

// 	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
// 		log.Fatalf("Migrate: up error: %s", err)
// 	}

// 	if errors.Is(err, migrate.ErrNoChange) {
// 		log.Printf("Migrate: no change")
// 		return
// 	}

// 	log.Printf("Migrate: up success")
// }
