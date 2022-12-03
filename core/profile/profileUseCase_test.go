package profile_test

import (
	"database/sql"
	"os"
	"proven/core/profile"
	"reflect"
	"testing"
	"time"

	database "proven/adapters/database"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

// fix relative paths
// test password hashing

var (
	db      *sql.DB
	useCase profile.ProfileUseCase
	m       *migrate.Migrate
)

func setup() {
	viper.SetConfigFile("../../config.test.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	connection := viper.GetString(`database.connection`)
	db, err = sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	m, err = migrate.New(
		"file://../../db/migrations",
		connection,
	)

	if err != nil {
		panic(err)
	}

	m.Up()

	profileRepo := database.NewProfileRepository(db)
	useCase = *profile.New(profileRepo)
}

func shutdown() {
	m.Down()
	db.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	input := profile.Profile{
		FirstName:        "Jamie",
		LastName:         "Foxx",
		Email:            "jamie@foxx.com",
		PhoneNumber:      "5124125",
		Citizenship:      "US",
		BirthDate:        time.Date(1988, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
		BirthCountry:     "US",
		ResidenceCountry: "US",
	}

	stored, err := useCase.Store(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := profile.Profile{
		FirstName:        stored.FirstName,
		LastName:         stored.LastName,
		Email:            stored.Email,
		PhoneNumber:      stored.PhoneNumber,
		Citizenship:      stored.Citizenship,
		BirthDate:        stored.BirthDate,
		BirthCountry:     stored.BirthCountry,
		ResidenceCountry: stored.ResidenceCountry,
	}
	if !reflect.DeepEqual(input, expected) {
		t.Error("Cannot store profile")
	}
}

func TestFetch(t *testing.T) {
	input := profile.Profile{
		Email: "jamie@foxx.com",
	}

	stored, err := useCase.Store(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	fetched, err := useCase.Get(stored.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	if input.Email != fetched.Email {
		t.Error("Cannot fetch profile")
	}
}

func TestUpdated(t *testing.T) {
	stored, _ := useCase.Store(profile.Profile{
		FirstName:        "Jamie",
		LastName:         "Foxx",
		Email:            "jamie@foxx.com",
		PhoneNumber:      "5124125",
		Citizenship:      "US",
		BirthDate:        time.Date(1988, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
		BirthCountry:     "US",
		ResidenceCountry: "US",
	})
	input := profile.Profile{
		FirstName:        "Hue",
		LastName:         "Jackman",
		PhoneNumber:      "1234567",
		Citizenship:      "AU",
		ResidenceCountry: "AU",
	}

	err := useCase.Update(stored.ID, input)
	if err != nil {
		t.Errorf(err.Error())
	}

	fetched, _ := useCase.Get(stored.ID)

	expected := profile.Profile{
		FirstName:        fetched.FirstName,
		LastName:         fetched.LastName,
		PhoneNumber:      fetched.PhoneNumber,
		Citizenship:      fetched.Citizenship,
		ResidenceCountry: fetched.ResidenceCountry,
	}
	if !reflect.DeepEqual(input, expected) {
		t.Error("Cannot update profile")
	}
}

func TestArchive(t *testing.T) {
	input := profile.Profile{
		Email: "jamie@foxx.com",
	}

	stored, err := useCase.Store(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = useCase.Archive(stored.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	archived, _ := useCase.Get(stored.ID)
	if archived.ArchivedAt.IsZero() {
		t.Errorf("Cannot archive profile")
	}
}
